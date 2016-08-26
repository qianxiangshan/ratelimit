//速率限制模块
//提供最大并发数的速率控制,
//提供单位时间请求数的控制(单位时间内请求书控制使用key标记,key不同则采用不同的令牌桶)
package ratelimit

import (
	//"fmt"
	"time"
)

//key   ratelimit key
// value lasttime rate/s timeinterval/nanosecond

type buket struct {
	//last call time
	Lasttime int64
	//interval
	Interval int64
}

type rateLimite struct {
	//key
	Key string
	//store buket
	Keybuket buket
}

// key 表示限制的对象,times 表示time的时间内允许的请求次数times/timeinterval
func newRateLimter(key string, times uint32, timeinterval time.Duration) *rateLimite {
	limiter := new(rateLimite)
	limiter.Key = key
	limiter.Keybuket.Lasttime = 0
	limiter.Keybuket.Interval = int64(timeinterval / time.Duration(times))
	return limiter
}

//判断请求是否允许接收
//依据最后一个访问,来限制访问速率
func (limiter *rateLimite) rateLimited() bool {
	currenttime := time.Now().UnixNano()
	interval := currenttime - limiter.Keybuket.Lasttime
	//fmt.Println(interval, limiter.Keybuket.Interval)
	if interval >= limiter.Keybuket.Interval {
		limiter.Keybuket.Lasttime = currenttime
		return true
	} else {
		return false
	}
}

func (limiter *rateLimite) adjustRateLimit(times uint32, timeinterval time.Duration) {
	limiter.Keybuket.Interval = int64(timeinterval / time.Duration(times))
}
