package cache

// KV 缓存系统
type Cache interface {
	// 初始化缓存服务
	Init(...Option) error
	// 获取键值
	Get(key string) (interface{}, error)
	// 新增、更新 键值
	Set(key string, val interface{}) error
	// 删除键
	Delete(key string) error
	// 返回具体实现名称
	String() string
}

// 配置项
type Options struct {
	Nodes []string
}

type Option func(o *Options)

// 设置连接的节点
func WithNodes(nodes ...string) Option {
	return func(o *Options) {
		o.Nodes = nodes
	}
}
