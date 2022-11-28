package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/yqchilde/pkgs/config"
)

var RClient *redis.Client

const (
	// ErrRedisNotFound redis为空
	ErrRedisNotFound = redis.Nil

	// DefaultRedisName 默认redis名称
	DefaultRedisName = "default"
)

// RManager redis管理器
type RManager struct {
	clients map[string]*redis.Client
	*sync.RWMutex
}

// NewRManager 创建redis管理器
func NewRManager() *RManager {
	return &RManager{
		clients: make(map[string]*redis.Client),
		RWMutex: &sync.RWMutex{},
	}
}

// Init 初始化redis
func Init() *redis.Client {
	clientManager := NewRManager()
	rdb, err := clientManager.GetClient(DefaultRedisName)
	if err != nil {
		panic(fmt.Sprintf("init redis client err: %s", err.Error()))
	}
	RClient = rdb
	return rdb
}

// GetClient 获取redis客户端
func (r *RManager) GetClient(name string) (*redis.Client, error) {
	// 读配置
	r.Lock()
	if client, ok := r.clients[name]; ok {
		r.RUnlock()
		return client, nil
	}
	r.RUnlock()
	c, err := LoadConf(name)
	if err != nil {
		panic(fmt.Sprintf("load redis config err: %v", err))
	}

	// 创建redis client
	r.Lock()
	defer r.Unlock()
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.DB,
		MinIdleConns: c.MinIdleConn,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		PoolSize:     c.PoolSize,
		PoolTimeout:  c.PoolTimeout,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	r.clients[name] = rdb
	return rdb, nil
}

// LoadConf 加载redis配置
func LoadConf(name string) (ret *Config, err error) {
	v, err := config.LoadWithType("redis", "yaml")
	if err != nil {
		return nil, err
	}

	var c Config
	err = v.UnmarshalKey(name, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// InitTestRedis 实例化一个可以用于单元测试的redis
func InitTestRedis() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// 打开下面命令可以测试链接关闭的情况
	// defer mr.Close()

	RClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	fmt.Println("mini redis addr:", mr.Addr())
}
