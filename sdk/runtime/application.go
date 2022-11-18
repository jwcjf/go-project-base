package runtime

import (
	"net/http"
	"sync"

	"github.com/jwcjf/go-project-base/logger"
	"github.com/jwcjf/go-project-base/storage"
	"github.com/jwcjf/go-project-base/storage/queue"
	"github.com/casbin/casbin/v2"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

// Application ...
type Application struct {
	dbs         map[string]*gorm.DB
	casbins     map[string]*casbin.SyncedEnforcer
	engine      http.Handler
	crontab     map[string]*cron.Cron
	mux         sync.RWMutex
	middlewares map[string]interface{}
	cache       storage.AdapterCache
	queue       storage.AdapterQueue
	locker      storage.AdapterLocker
	memoryQueue storage.AdapterQueue
}

// SetDb 设置对应key的db
func (e *Application) SetDb(key string, db *gorm.DB) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.dbs[key] = db
}

// GetDb 获取所有map里的db数据
func (e *Application) GetDb() map[string]*gorm.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.dbs
}

// GetDbByKey 根据key获取db
func (e *Application) GetDbByKey(key string) *gorm.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	if db, ok := e.dbs["*"]; ok {
		return db
	}
	return e.dbs[key]
}

// SetCasbin ...
func (e *Application) SetCasbin(key string, enforcer *casbin.SyncedEnforcer) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.casbins[key] = enforcer
}

// GetCasbin ...
func (e *Application) GetCasbin() map[string]*casbin.SyncedEnforcer {
	return e.casbins
}

// GetCasbinKey 根据key获取casbin
func (e *Application) GetCasbinKey(key string) *casbin.SyncedEnforcer {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e, ok := e.casbins["*"]; ok {
		return e
	}
	return e.casbins[key]
}

// SetEngine 设置路由引擎
func (e *Application) SetEngine(engine http.Handler) {
	e.engine = engine
}

// GetEngine 获取路由引擎
func (e *Application) GetEngine() http.Handler {
	return e.engine
}

// SetLogger 设置日志组件
func (e *Application) SetLogger(l logger.Logger) {
	logger.DefaultLogger = l
}

// GetLogger 获取日志组件
func (e *Application) GetLogger() logger.Logger {
	return logger.DefaultLogger
}

// NewConfig 默认值
func NewConfig() *Application {
	return &Application{
		dbs:         make(map[string]*gorm.DB),
		casbins:     make(map[string]*casbin.SyncedEnforcer),
		crontab:     make(map[string]*cron.Cron),
		middlewares: make(map[string]interface{}),
		memoryQueue: queue.NewMemory(10000),
	}
}

// SetCrontab 设置对应key的crontab
func (e *Application) SetCrontab(key string, crontab *cron.Cron) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.crontab[key] = crontab
}

// GetCrontab 获取所有map里的crontab数据
func (e *Application) GetCrontab() map[string]*cron.Cron {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.crontab
}

// GetCrontabKey 根据key获取crontab
func (e *Application) GetCrontabKey(key string) *cron.Cron {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e, ok := e.crontab["*"]; ok {
		return e
	}
	return e.crontab[key]
}

// SetMiddleware 设置中间件
func (e *Application) SetMiddleware(key string, middleware interface{}) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.middlewares[key] = middleware
}

// GetMiddleware 获取所有中间件
func (e *Application) GetMiddleware() map[string]interface{} {
	return e.middlewares
}

// GetMiddlewareKey 获取对应key的中间件
func (e *Application) GetMiddlewareKey(key string) interface{} {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.middlewares[key]
}

// SetCacheAdapter 设置缓存
func (e *Application) SetCacheAdapter(c storage.AdapterCache) {
	e.cache = c
}

// GetCacheAdapter 获取缓存
func (e *Application) GetCacheAdapter() storage.AdapterCache {
	return NewCache("", e.cache, "")
}

// GetCachePrefix 获取带租户标记的cache
func (e *Application) GetCachePrefix(key string) storage.AdapterCache {
	return NewCache(key, e.cache, "")
}

// SetQueueAdapter 设置队列适配器
func (e *Application) SetQueueAdapter(c storage.AdapterQueue) {
	e.queue = c
}

// GetQueueAdapter 获取队列适配器
func (e *Application) GetQueueAdapter() storage.AdapterQueue {
	return NewQueue("", e.queue)
}

// GetQueuePrefix 获取带租户标记的queue
func (e *Application) GetQueuePrefix(key string) storage.AdapterQueue {
	return NewQueue(key, e.queue)
}

// SetLockerAdapter 设置分布式锁
func (e *Application) SetLockerAdapter(c storage.AdapterLocker) {
	e.locker = c
}

// GetLockerAdapter 获取分布式锁
func (e *Application) GetLockerAdapter() storage.AdapterLocker {
	return NewLocker("", e.locker)
}

// GetLockerPrefix ...
func (e *Application) GetLockerPrefix(key string) storage.AdapterLocker {
	return NewLocker(key, e.locker)
}

// GetStreamMessage 获取队列需要用的message
func (e *Application) GetStreamMessage(id, stream string, value map[string]interface{}) (storage.Messager, error) {
	message := &queue.Message{}
	message.SetID(id)
	message.SetStream(stream)
	message.SetValues(value)
	return message, nil
}

// GetMemoryQueue ...
func (e *Application) GetMemoryQueue(prefix string) storage.AdapterQueue {
	return NewQueue(prefix, e.memoryQueue)
}
