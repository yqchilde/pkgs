package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/yqchilde/gint/pkg/logger"
)

var Cfg = &Config{}

const (
	configDefaultPath = "config"
	configDefaultName = "config"
)

// Init config
func Init(configPath ...string) (*Config, error) {
	var path string
	if len(configPath) == 1 {
		path = configPath[0]
	}
	cfgFile, err := LoadConfig(path)
	if err != nil {
		logger.Fatalf("LoadConfig: %v", err)
	}

	Cfg, err = ParseConfig(cfgFile)
	if err != nil {
		logger.Fatalf("ParseConfig: %v", err)
	}

	go watchConfig(cfgFile)

	return Cfg, nil
}

// LoadConfig load config file from given path
func LoadConfig(confPath string) (*viper.Viper, error) {
	v := viper.New()
	if confPath != "" {
		// 指定路径
		v.SetConfigFile(confPath)
	} else {
		// 默认路径
		v.AddConfigPath(configDefaultPath)
		v.SetConfigName(configDefaultName)
	}
	v.SetConfigType("yaml")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %s", err.Error())
	}

	return &c, nil
}

// watchConfig
func watchConfig(v *viper.Viper) {
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("[Config] file changed: %s", e.Name)
	})
}
