package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	conf         *Config
	FileTypeYaml = "yaml"
	FileTypeJson = "json"
	FileTypeToml = "toml"
)

type Config struct {
	env        string
	configDir  string
	configType string
	val        map[string]*viper.Viper
	mu         sync.Mutex
}

// New 创建一个config实例
func New(cfgDir string, opts ...Option) *Config {
	if cfgDir == "" {
		panic("config dir is not set")
	}
	c := Config{
		configDir:  cfgDir,
		configType: FileTypeYaml,
		val:        make(map[string]*viper.Viper),
	}
	for _, opt := range opts {
		opt(&c)
	}
	conf = &c
	return &c
}

func Load(filename string, val interface{}) error { return conf.Load(filename, val) }

func (c *Config) Load(filename string, val interface{}) error {
	v, err := c.LoadWithType(filename, c.configType)
	if err != nil {
		return err
	}
	if err := v.Unmarshal(&val); err != nil {
		return err
	}
	return nil
}

func LoadYaml(filename string, val interface{}) error { return conf.LoadYaml(filename, val) }

func (c *Config) LoadYaml(filename string, val interface{}) error {
	v, err := c.LoadWithType(filename, FileTypeYaml)
	if err != nil {
		return err
	}
	err = v.Unmarshal(&val)
	if err != nil {
		return err
	}
	return nil
}

// LoadJson alias for config func.
func LoadJson(filename string, val interface{}) error { return conf.LoadJson(filename, val) }

// LoadJson scan data to struct.
func (c *Config) LoadJson(filename string, val interface{}) error {
	v, err := c.LoadWithType(filename, FileTypeJson)
	if err != nil {
		return err
	}
	err = v.Unmarshal(&val)
	if err != nil {
		return err
	}
	return nil
}

// LoadToml alias for config func.
func LoadToml(filename string, val interface{}) error { return conf.LoadToml(filename, val) }

// LoadToml scan data to struct.
func (c *Config) LoadToml(filename string, val interface{}) error {
	v, err := c.LoadWithType(filename, FileTypeToml)
	if err != nil {
		return err
	}
	err = v.Unmarshal(&val)
	if err != nil {
		return err
	}
	return nil
}

// LoadWithType 加载配置文件
func LoadWithType(filename string, cfgType string) (v *viper.Viper, err error) {
	return conf.LoadWithType(filename, cfgType)
}

func (c *Config) LoadWithType(filename string, cfgType string) (v *viper.Viper, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.val[filename]; ok {
		return v, nil
	}
	v, err = c.load(filename, cfgType)
	if err != nil {
		return nil, err
	}
	c.val[filename] = v
	return v, nil
}

func (c *Config) load(filename string, cfgType string) (*viper.Viper, error) {
	env := GetEnvString("APP_ENV", "")
	path := filepath.Join(c.configDir, env)
	if c.env != "" {
		path = filepath.Join(c.configDir, c.env)
	}
	if cfgType != "" {
		c.configType = cfgType
	}

	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(filename)
	v.SetConfigType(c.configType)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	// 监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file changed: %s", e.Name)
	})
	return v, nil
}

func GetEnvString(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}
