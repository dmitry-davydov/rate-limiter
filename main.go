package main

import (
	"fmt"
	"sync/atomic"
	"time"
)



func main() {
	limiter := NewLimiter(100)

	var counter int32 = 0

	go func (cntptr *int32){
		ticker := time.NewTicker(time.Second)
		for {
			t := <- ticker.C
			fmt.Printf("%s: cnt: %d\n", t, atomic.LoadInt32(cntptr))
			atomic.StoreInt32(cntptr, 0)
		}
	}(&counter)
	
	for {
		if !limiter.Allow() {

		} else {
			atomic.AddInt32(&counter, 1)
		}
	}
}
