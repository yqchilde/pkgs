package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

var (
	RC         *redis.Client
	ValueTrue  = 1
	ValueFalse = 0
)

// ErrRedisNotFound not exist in redis
const ErrRedisNotFound = redis.Nil

// Config for redis base config
type Config struct {
	Addr         string        `mapstructure:"addr"`
	Password     string        `mapstructure:"password"`
	DB           int           `mapstructure:"db"`
	MinIdleConn  int           `mapstructure:"min-idle-conn"`
	PoolSize     int           `mapstructure:"pool-size"`
	IsTrace      bool          `mapstructure:"is-trace"`
	DialTimeout  time.Duration `mapstructure:"dial-timeout"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout"`
	PoolTimeout  time.Duration `mapstructure:"pool-timeout"`
}

// Init for init redis client
func Init(c *Config) *redis.Client {
	RC = redis.NewClient(&redis.Options{
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

	_, err := RC.Ping(context.Background()).Result()
	if err != nil {
		log.Panicf("[redis] redis ping err: %+v", err)
	}

	// hook tracing (with open telemetry)
	if c.IsTrace {
		RC.AddHook(redisotel.NewTracingHook())
	}

	return RC
}

// InitTestRedis A redis instance for unit testing
func InitTestRedis() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	// defer mr.Close()
	RC = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	fmt.Println("mini redis addr:", mr.Addr())
}
