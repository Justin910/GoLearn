package main

import (
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

func main123() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var str string
	var at atomic.Value
	go func() {
		for {
			at.Store(str)

			log.Printf("%s\n", at.Load().(string))
			//fmt.Sprintf("%s\n", str)
			//a := str
			//_ = a
			//log.Printf("%s\n", a)
		}
	}()

	for {
		str = ""
		time.Sleep(10 * time.Nanosecond)
		str = "/test/test/test"
		time.Sleep(10 * time.Nanosecond)
	}
}
