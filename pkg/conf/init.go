package conf

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/yqchilde/gin-skeleton/pkg/env"
)

var Conf = &Config{}

const configDefaultPath = "config"

// Init config
func Init() (*Config, error) {
	cfgFile, err := LoadConfig(configDefaultPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	Conf, err = ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	go watchConfig(cfgFile)

	return Conf, nil
}

// LoadConfig load config file from given path
func LoadConfig(confPath string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(confPath)
	v.SetConfigName(fmt.Sprintf("config.%s", env.GetEnv()))
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
		log.Printf("Config file changed: %s", e.Name)
	})
}
