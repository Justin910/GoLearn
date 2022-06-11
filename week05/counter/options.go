package counter

import "time"

type Option func(o *options)

// options is an application options.
type options struct {
	bucketWidth time.Duration
	windowSize  int32
}

// WithBucketWidth
// 设置桶宽度, 默认 DefaultBucketWidth
func WithBucketWidth(w time.Duration) Option {
	return func(o *options) {
		o.bucketWidth = w
	}
}

// WithWindowSize
// 滑动窗口大小, 默认 DefaultWindowSize
func WithWindowSize(size int32) Option {
	return func(o *options) {
		o.windowSize = size
	}
}
