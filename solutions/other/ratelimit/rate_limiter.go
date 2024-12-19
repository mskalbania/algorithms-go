package ratelimit

import (
	"context"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(ctx context.Context, limitPerMin int) *RateLimiter {
	r := &RateLimiter{make(chan struct{}, limitPerMin)}
	go r.tokenFiller(ctx, limitPerMin)
	return r
}

func (r *RateLimiter) tokenFiller(ctx context.Context, limitPerMin int) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for i := 0; i < limitPerMin; i++ {
				select {
				case r.tokens <- struct{}{}: //based on idea of tokens in bucket that are refilled at const interval
				default:
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (r *RateLimiter) Allow() bool {
	select {
	case <-r.tokens:
		return true
	default:
		return false
	}
}

func (r *RateLimiter) Wait() {
	<-r.tokens
}
