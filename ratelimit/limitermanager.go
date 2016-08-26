package ratelimit

import (
	"dana-tech.com/pg/common/cache"
	"time"
)

type RateLimiters struct {
	ratelimiters *cache.MemCache
	Interval     int64
}

// keycounts 表示最大限制数的key,多于最大数的key会被按照lru算法移除,未存在的可以则重新记录,rate指该类的key每个都是这个速率
func NewRateLimters(keycounts int, times uint32, timeinterval time.Duration) *RateLimiters {
	limiters := new(RateLimiters)
	limiters.ratelimiters = cache.NewMemCache(keycounts)
	limiters.Interval = int64(timeinterval / time.Duration(times))
	return limiters
}

//是否允许该key通过
func (limiters *RateLimiters) RateLimited(key string) bool {
	value, has := limiters.ratelimiters.Get(key)
	if has {
		//存在key, 计算是否可以通过
		limiter := value.(*rateLimite)
		return limiter.rateLimited()
	} else {
		//不存在key,则需要构建出key
		var limiter rateLimite
		limiter.Keybuket.Interval = limiters.Interval
		limiter.Key = key
		limiter.Keybuket.Lasttime = time.Now().UnixNano()
		//设置到cache中
		limiters.ratelimiters.Set(key, &limiter)
		return true
	}
}
