package runtime

import (
	"github.com/bsm/redislock"

	"git.code.oa.com/ad/go-project-base/storage"
)

// NewLocker 创建对应上下文分布式锁
func NewLocker(prefix string, locker storage.AdapterLocker) storage.AdapterLocker {
	return &Locker{
		prefix: prefix,
		locker: locker,
	}
}

// Locker ...
type Locker struct {
	prefix string
	locker storage.AdapterLocker
}

// String ...
func (e *Locker) String() string {
	return e.locker.String()
}

// Lock 返回分布式锁对象
func (e *Locker) Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error) {
	return e.locker.Lock(e.prefix+intervalTenant+key, ttl, options)
}
