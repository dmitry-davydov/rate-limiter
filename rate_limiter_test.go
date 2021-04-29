package main

import (
	"testing"
)

func Benchmark10rps(t *testing.B) {
	limiter := NewLimiter(100)
	for i := 0; i < t.N; i++ {
		limiter.UpdateLastRequest()
		limiter.Wait()
	}
}
