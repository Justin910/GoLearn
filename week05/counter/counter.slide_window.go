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

	DefaultErrPerceng = 0.2
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

	statCounter BucketCounterStream

	isStopCounter int32

	closeOnce sync.Once

	mux sync.RWMutex // calc lock
}

func InitSlideWindowCounter(opts ...Option) (*SlideWindow, error) {

	o := options{
		bucketWidth: DefaultBucketWidth,
		windowSize:  DefaultWindowSize,
		errPercent:  DefaultErrPerceng,
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

	if o.errPercent < 0.0 || o.errPercent > 1.0 {
		return nil, errors.New("The err percent param invalid ")
	}

	rollw := new(SlideWindow)
	rollw.offset = 0
	rollw.buckets = make([]BucketCounterStream, o.windowSize)
	rollw.opts = o

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
	r.mux.RLock()
	defer r.mux.RUnlock()

	counter := BucketCounterStream{
		Success:   atomic.LoadInt32(&r.statCounter.Success) + atomic.LoadInt32(&r.buckets[r.offset].Success),
		Failure:   atomic.LoadInt32(&r.statCounter.Failure) + atomic.LoadInt32(&r.buckets[r.offset].Failure),
		Timeout:   atomic.LoadInt32(&r.statCounter.Timeout) + atomic.LoadInt32(&r.buckets[r.offset].Timeout),
		Rejection: atomic.LoadInt32(&r.statCounter.Rejection) + atomic.LoadInt32(&r.buckets[r.offset].Rejection),
	}

	return counter
}

func (r *SlideWindow) IsHealthy() bool {
	cc := r.GetCurrentCounter()
	ep := cc.GetErrPercent()
	return ep < r.opts.errPercent
}

func (bcs BucketCounterStream) Sum() int32 {
	return atomic.LoadInt32(&bcs.Success) + atomic.LoadInt32(&bcs.Failure) + atomic.LoadInt32(&bcs.Timeout) + atomic.LoadInt32(&bcs.Rejection)
}

func (bcs BucketCounterStream) GetErrPercent() float64 {
	sum := bcs.Sum()
	if sum == 0 {
		return 0
	}
	return 1 - float64(atomic.LoadInt32(&bcs.Success))/float64(sum)
}

func (r *SlideWindow) Stop() {
	r.closeOnce.Do(func() {
		atomic.StoreInt32(&r.isStopCounter, 1)
		r.cancel()
	})
}

func (r *SlideWindow) Calc() {
	r.mux.Lock()
	defer r.mux.Unlock()

	// 调整偏移
	r.adjustOffset()

	// 计算窗口内的数据
	r.calc()
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

	r.mux.RLock()
	defer r.mux.RUnlock()

	switch ct {
	case counterType_Success:
		atomic.AddInt32(&r.buckets[r.offset].Success, n)
	case counterType_Failure:
		atomic.AddInt32(&r.buckets[r.offset].Failure, n)
	case counterType_Timeout:
		atomic.AddInt32(&r.buckets[r.offset].Timeout, n)
	case counterType_Rejection:
		atomic.AddInt32(&r.buckets[r.offset].Rejection, n)
	}
	return
}

func (r *SlideWindow) start() {

	timer := time.NewTicker(r.opts.bucketWidth)

	for {
		select {
		case <-r.ctx.Done():
			goto skipLoop

		case <-timer.C:
			r.Calc()
		}
	}
skipLoop:

	timer.Stop()
}

func (r *SlideWindow) calc() {
	statCounter := BucketCounterStream{}

	for i := range r.buckets {
		if int32(i) == r.offset {
			// 当前偏移窗口实时计算
			continue
		}
		atomic.AddInt32(&statCounter.Success, atomic.LoadInt32(&r.buckets[i].Success))
		atomic.AddInt32(&statCounter.Failure, atomic.LoadInt32(&r.buckets[i].Failure))
		atomic.AddInt32(&statCounter.Timeout, atomic.LoadInt32(&r.buckets[i].Timeout))
		atomic.AddInt32(&statCounter.Rejection, atomic.LoadInt32(&r.buckets[i].Rejection))
	}
	r.statCounter = statCounter
}

func (r *SlideWindow) adjustOffset() {
	// 计数桶偏移
	r.offset = atomic.AddInt32(&r.offset, 1) % r.opts.windowSize
	// 桶归零
	r.buckets[r.offset] = BucketCounterStream{}
}
