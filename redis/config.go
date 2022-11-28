package redis

import "time"

type Config struct {
	Addr         string        // redis地址
	Password     string        // redis密码
	DB           int           // redis数据库
	MinIdleConn  int           // 最小空闲连接数
	DialTimeout  time.Duration // 连接超时
	ReadTimeout  time.Duration // 读超时
	WriteTimeout time.Duration // 写超时
	PoolSize     int           // 连接池大小
	PoolTimeout  time.Duration // 连接池超时
}
