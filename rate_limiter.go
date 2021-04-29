package main

import (
	"time"
)

type Limiter struct {
	tick        time.Duration
	lastRequest time.Time
}

func NewLimiter(rps int64) *Limiter {
	limiter := new(Limiter)
	limiter.tick = time.Second / time.Duration(rps)
	limiter.lastRequest = time.Unix(0, 0)
	return limiter
}

func (t *Limiter) Wait() time.Duration {
	dx := time.Now().Sub(t.lastRequest)
	if dx >= t.tick {
		return 0
	}

	return t.tick - dx
}

func (t *Limiter) UpdateLastRequest() {
	t.lastRequest = time.Now()
}

func (t *Limiter) Allow() bool {
	if t.Wait() > 0 {
		return false
	}
	t.UpdateLastRequest()
	return true
}