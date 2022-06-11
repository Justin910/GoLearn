package counter

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// DefaultBucketWidth : 每个桶宽度
	DefaultBucketWidth = time.Millisecond * 100

	// DefaultWindowSize : 滑动窗口大小
	DefaultWindowSize = 10
)

// BucketCounterStream
// 滑动窗口计数
type BucketCounterStream struct {
	Success   int32
	Failure   int32
	Timeout   int32
	Rejection int32
}

type SlideWindow struct {
	ctx    context.Context
	cancel context.CancelFunc

	opts options

	offset  int32                 //桶偏移
	buckets []BucketCounterStream //桶数组

	ch          chan struTask
	statCounter BucketCounterStream

	isStopCounter int32

	closeOnce sync.Once

	sendWg sync.WaitGroup

	mux sync.Mutex // calc lock
}

func InitSlideWindowCounter(opts ...Option) (*SlideWindow, error) {

	o := options{
		bucketWidth: DefaultBucketWidth,
		windowSize:  DefaultWindowSize,
	}

	for _, opt := range opts {
		opt(&o)
	}

	if o.bucketWidth < time.Millisecond*10 {
		return nil, errors.New("Bucket width too short ")
	}

	if o.windowSize == 0 {
		return nil, errors.New("Window size is 0 ")
	}

	if o.windowSize > 1000 {
		return nil, errors.New("The window size is too large ")
	}

	rollw := new(SlideWindow)
	rollw.offset = 0
	rollw.buckets = make([]BucketCounterStream, o.windowSize)
	rollw.opts = o

	rollw.ch = make(chan struTask, 1000)
	rollw.ctx, rollw.cancel = context.WithCancel(context.Background())

	go rollw.start()
	return rollw, nil
}

func (r *SlideWindow) IncSuccess(n int32) {
	r.inc(counterType_Success, n)
}
func (r *SlideWindow) IncFailure(n int32) {
	r.inc(counterType_Failure, n)
}
func (r *SlideWindow) IncTimeout(n int32) {
	r.inc(counterType_Timeout, n)
}
func (r *SlideWindow) IncRejection(n int32) {
	r.inc(counterType_Rejection, n)
}

func (r *SlideWindow) GetCurrentCounter() BucketCounterStream {
	return r.statCounter
}

func (r *SlideWindow) Stop() {
	r.closeOnce.Do(func() {
		atomic.StoreInt32(&r.isStopCounter, 1)
		r.cancel()
		r.sendWg.Wait()
		close(r.ch)
	})
}

func (r *SlideWindow) Calc() {
	r.mux.Lock()
	defer r.mux.Unlock()

	// 计算窗口内的数据
	r.calc()
	// 调整偏移
	r.adjustOffset()
}

type counterType int

const (
	counterType_Success counterType = 1 << iota
	counterType_Failure
	counterType_Timeout
	counterType_Rejection
)

type struTask struct {
	ct counterType
	n  int32
}

func (r *SlideWindow) inc(ct counterType, n int32) {

	// 避免已经stop后，还对wait.group进行添加任务
	if atomic.LoadInt32(&r.isStopCounter) == 1 {
		return
	}

	r.sendWg.Add(1)
	defer r.sendWg.Done()

	// 避免在刚判断完是否stop后，还没走到Add方法时，Stop函数就已经调用.Wait()了
	if atomic.LoadInt32(&r.isStopCounter) == 1 {
		return
	}

	t := struTask{
		ct: ct,
		n:  n,
	}

	select {
	case <-r.ctx.Done():
	case r.ch <- t:
	default:
	}
}

func (r *SlideWindow) start() {

	timer := time.NewTicker(r.opts.bucketWidth)

	for {
		select {
		case <-r.ctx.Done():
			goto skipLoop

		case t := <-r.ch:
			// 单goroutine处理，不需要用原子锁
			switch t.ct {
			case counterType_Success:
				r.buckets[r.offset].Success += t.n
			case counterType_Failure:
				r.buckets[r.offset].Failure += t.n
			case counterType_Timeout:
				r.buckets[r.offset].Timeout += t.n
			case counterType_Rejection:
				r.buckets[r.offset].Rejection += t.n
			}

		case <-timer.C:
			r.Calc()
		}
	}
skipLoop:

	timer.Stop()

	for range r.ch {
		// Discard Message
	}
}

func (r *SlideWindow) calc() {
	statCounter := BucketCounterStream{}

	for i := range r.buckets {
		statCounter.Failure += r.buckets[i].Failure
		statCounter.Success += r.buckets[i].Success
		statCounter.Timeout += r.buckets[i].Timeout
		statCounter.Rejection += r.buckets[i].Rejection
	}
	r.statCounter = statCounter
}

func (r *SlideWindow) adjustOffset() {
	// 计数桶偏移
	r.offset = atomic.AddInt32(&r.offset, 1) % r.opts.windowSize
	// 桶归零
	r.buckets[r.offset] = BucketCounterStream{}
}
