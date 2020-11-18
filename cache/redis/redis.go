package redis

import (
	"github.com/gomodule/redigo/redis"
	redisCluster "github.com/mna/redisc"
	"kit/cache"
	"log"
	"time"
)

// redis实现的缓存系统
type redisCache struct {
	options cache.Options        // 缓存配置
	client  redisCluster.Cluster // redis 集群客户端
}

// 可以在这里初始化配置项
func (rc *redisCache) Init(opts ...cache.Option) error {
	for _, o := range opts {
		o(&rc.options)
	}
	return nil
}

func (rc *redisCache) Get(key string) (interface{}, error) {
	// 获取redis客户度，并发安全的
	conn := rc.client.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

func (rc *redisCache) Set(key string, val interface{}) error {
	conn := rc.client.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	return err
}

func (rc *redisCache) Delete(key string) error {
	conn := rc.client.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

func (rc *redisCache) String() string {
	return "redis"
}

func NewCache(opts ...cache.Option) cache.Cache {
	// 初始化配置项
	var options cache.Options

	// 如果有，可以在这里先设置默认值
	options.Nodes = []string{"127.0.0.1:6379", "127.0.0.1:6380", "127.0.0.1:6381"}

	// 这里可以覆盖前面的默认值
	for _, o := range opts {
		o(&options)
	}

	// 初始化redis集群
	cluster := redisCluster.Cluster{
		StartupNodes: options.Nodes,
		DialOptions:  []redis.DialOption{redis.DialConnectTimeout(5 * time.Second)},
		CreatePool:   createPool,
		//PoolWaitTime: 0,
	}

	if err := cluster.Refresh(); err != nil {
		log.Fatalf("NewCache cluster.Refresh error(%v)", err)
	}

	return &redisCache{
		options: options,
		client:  cluster,
	}
}

func createPool(address string, options ...redis.DialOption) (*redis.Pool, error) {
	pool := redis.Pool{
		MaxIdle:     5,
		MaxActive:   10,
		IdleTimeout: time.Millisecond * 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &pool, nil
}
