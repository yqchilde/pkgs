package config

import "time"

// APIConfig for http and rpc config
type APIConfig struct {
	Addr         string        `json:"addr"`
	ReadTimeout  time.Duration `json:"read-timeout"`
	WriteTimeout time.Duration `json:"write-timeout"`
}
