package locker

import (
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v7"
)

// NewRedis 初始化locker
func NewRedis(c *redis.Client) *Redis {
	return &Redis{
		client: c,
	}
}

// Redis ...
type Redis struct {
	client *redis.Client
	mutex  *redislock.Client
}

// String ...
func (Redis) String() string {
	return "redis"
}

// Lock ...
func (r *Redis) Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error) {
	if r.mutex == nil {
		r.mutex = redislock.New(r.client)
	}
	return r.mutex.Obtain(key, time.Duration(ttl)*time.Second, options)
}
