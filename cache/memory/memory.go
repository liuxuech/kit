package memory

import (
	"errors"
	"kit/cache"
	"sync"
)

// 缓存的内存实现
type memoryCache struct {
	sync.RWMutex                        // 并发控制
	values       map[string]interface{} // 数据存储位置
}

func (m *memoryCache) Init(...cache.Option) error {
	// 这里不需要初始化
	return nil
}

func (m *memoryCache) Get(key string) (interface{}, error) {
	// 并发控制，读锁
	m.RLock()
	defer m.RUnlock()

	if v, ok := m.values[key]; !ok {
		return nil, errors.New(key + " 不存在")
	} else {
		return v, nil
	}
}

func (m *memoryCache) Set(key string, val interface{}) error {
	m.Lock()
	m.values[key] = val
	m.Unlock()
	return nil
}

func (m *memoryCache) Delete(key string) error {
	m.Lock()
	delete(m.values, key)
	m.Unlock()
	return nil
}

func (m *memoryCache) String() string {
	return "内存"
}

// 返回缓存接口
func NewCache(opts ...cache.Option) cache.Cache {
	return &memoryCache{
		values: make(map[string]interface{}),
	}
}
