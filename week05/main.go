package main

import (
	"context"
	"fmt"
	"time"
	"week05/counter"
)

const (
	// BucketWidth : 每个通宽度为1秒
	BucketWidth = time.Second

	// WindowSize : 滑动窗口大小
	WindowSize = 10
)

func main() {

	r, err := counter.InitSlideWindowCounter(counter.WithBucketWidth(time.Millisecond*100), counter.WithWindowSize(30))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {

			select {
			case <-ctx.Done():
				return
			default:
			}
			r.IncSuccess(1)
			r.IncFailure(2)
			r.IncTimeout(3)
			r.IncRejection(4)
		}
	}()

	num := 0
	for {
		num++
		if num == 5 {
			break
		}
		cc := r.GetCurrentCounter()
		fmt.Println(cc)
		fmt.Println(cc.GetErrPercent())
		time.Sleep(time.Second)
	}

	r.Stop()
	time.Sleep(time.Second * 2)
}
