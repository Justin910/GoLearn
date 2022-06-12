package counter

import (
	"testing"
	"time"
)

func TestSlideWindow_SingleWindow(t *testing.T) {

	success := int32(1)
	failure := int32(2)
	timeout := int32(3)
	rejection := int32(4)

	size := int32(1)
	bucketWidth := time.Second * 10 // 延长自动计算间隔，

	sw, _ := InitSlideWindowCounter(WithWindowSize(size), WithBucketWidth(bucketWidth))
	defer sw.Stop()

	sw.IncSuccess(success)
	sw.IncFailure(failure)
	sw.IncTimeout(timeout)
	sw.IncRejection(rejection)

	cc := sw.GetCurrentCounter()

	if cc.Success != success {
		t.Fatal("Counter Success Error")
	}

	if cc.Failure != failure {
		t.Fatal("Counter Failure Error")
	}

	if cc.Timeout != timeout {
		t.Fatal("Counter Timeout Error")
	}

	if cc.Rejection != rejection {
		t.Fatal("Counter Rejection Error")
	}
}

func TestSlideWindow_MultiWindow(t *testing.T) {

	success := int32(1)
	failure := int32(2)
	timeout := int32(3)
	rejection := int32(4)

	size := int32(3)
	bucketWidth := time.Second * 1000 // 延长自动计算间隔，采用手动计算方式

	sw, _ := InitSlideWindowCounter(WithWindowSize(size), WithBucketWidth(bucketWidth))
	defer sw.Stop()

	for i := int32(0); i < size; i++ {
		sw.IncSuccess(success)
		sw.IncFailure(failure)
		sw.IncTimeout(timeout)
		sw.IncRejection(rejection)

		if i != size-1 {
			sw.Calc()
		}
	}

	for _, c := range sw.buckets {
		if c.Success != success {
			t.Fatal("Counter Success Error")
		}

		if c.Failure != failure {
			t.Fatal("Counter Failure Error")
		}

		if c.Timeout != timeout {
			t.Fatal("Counter Timeout Error")
		}

		if c.Rejection != rejection {
			t.Fatal("Counter Rejection Error")
		}
	}

	cc := sw.GetCurrentCounter()

	if cc.Success != success*size {
		t.Fatal("Counter Success Error")
	}

	if cc.Failure != failure*size {
		t.Fatal("Counter Failure Error")
	}

	if cc.Timeout != timeout*size {
		t.Fatal("Counter Timeout Error")
	}

	if cc.Rejection != rejection*size {
		t.Fatal("Counter Rejection Error")
	}
}

func TestSlideWindow_ErrPercent_NoHealthy(t *testing.T) {

	success := int32(1)
	failure := int32(2)
	timeout := int32(3)
	rejection := int32(4)

	size := int32(2)
	bucketWidth := time.Second * 10 // 延长自动计算间隔，

	sw, _ := InitSlideWindowCounter(WithWindowSize(size), WithBucketWidth(bucketWidth), WithErrPercent(0.2))
	defer sw.Stop()

	for i := int32(0); i < size; i++ {
		sw.IncSuccess(success)
		sw.IncFailure(failure)
		sw.IncTimeout(timeout)
		sw.IncRejection(rejection)

		if i != size-1 {
			sw.Calc()
		}
	}

	if sw.IsHealthy() {
		t.Fatal("Check Healthy Error")
	}
}

func TestSlideWindow_ErrPercent_Healthy(t *testing.T) {

	success := int32(20)
	failure := int32(1)
	timeout := int32(1)
	rejection := int32(1)

	size := int32(2)
	bucketWidth := time.Second * 10 // 延长自动计算间隔，

	sw, _ := InitSlideWindowCounter(WithWindowSize(size), WithBucketWidth(bucketWidth), WithErrPercent(0.2))
	defer sw.Stop()

	for i := int32(0); i < size; i++ {
		sw.IncSuccess(success)
		sw.IncFailure(failure)
		sw.IncTimeout(timeout)
		sw.IncRejection(rejection)

		if i != size-1 {
			sw.Calc()
		}
	}

	if !sw.IsHealthy() {
		t.Fatal("Check Healthy Error")
	}
}
