package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRateLimters(t *testing.T) {
	var r *RateLimiters
	r = NewRateLimters(10, 1, 3*time.Millisecond)
	if r.ratelimiters.Status().CurrentSize != 0 {
		t.Fail()
	}
	if r.ratelimiters.Status().MaxItemSize != 10 {
		t.Fail()
	}
}

func TestRateLimiteds(t *testing.T) {
	var r *RateLimiters
	r = NewRateLimters(2, 1, 25*time.Millisecond)

	fmt.Println(r.RateLimited("gg"))
	//time.Sleep(time.Millisecond * 3)
	fmt.Println(r.RateLimited("gg"))
	fmt.Println(r.RateLimited("test"))
	time.Sleep(time.Millisecond * 50)
	fmt.Println(r.RateLimited("test"))
	fmt.Println(r.ratelimiters.Status())

	fmt.Println(r.RateLimited("ggg"))
	fmt.Println(r.ratelimiters.Status())

	//	if true != r.rateLimited() {
	//		t.Fail()
	//	}
	//	//time.Sleep(time.Microsecond)
	//	if true != r.rateLimited() {
	//		t.Fail()
	//	}
}

var rs = NewRateLimters(10, 1, 3*time.Millisecond)

func BenchmarkRateLimiteds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(rs.RateLimited("gg"))
	}
}
