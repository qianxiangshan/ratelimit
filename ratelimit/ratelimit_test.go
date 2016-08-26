package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRateLimter(t *testing.T) {
	var r *rateLimite
	r = newRateLimter("test", 1, 3*time.Millisecond)
	if r.Keybuket.Interval != 3000000 {
		t.Fail()
	}
}

func TestRateLimited(t *testing.T) {
	var r *rateLimite
	r = newRateLimter("test", 1, time.Microsecond*25)

	fmt.Println(r.rateLimited())
	fmt.Println(r.rateLimited())

	//	if true != r.rateLimited() {
	//		t.Fail()
	//	}
	//	//time.Sleep(time.Microsecond)
	//	if true != r.rateLimited() {
	//		t.Fail()
	//	}
}

var r = newRateLimter("test", 1, time.Millisecond*100)

func BenchmarkRateLimited(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(r.rateLimited())
	}
}

func TestadjustRateLimit(t *testing.T) {

}
