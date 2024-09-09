package config

import (
	"log"
	"os"
	"sync"

	yaml "gopkg.in/yaml.v3"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Port     int      `yaml:"port"`
	Debug    bool     `yaml:"debug"`
	Database database `yaml:"database"`
	Redis    redis    `yaml:"redis"`
	Chatgpt  chatgpt  `yaml:"chatgpt"`
}

type database struct {
	DataSourceName        string  `yaml:"data_source_name"`
	GormCfg               gormCfg `yaml:"gorm_cfg"`
	MaxOpenConnections    int     `yaml:"max_open_connections"`
	ConnectionMaxLifetime int     `yaml:"connection_max_lifetime"`
}

type redis struct {
	Addr   string `yaml:"addr"`
	Passwd string `yaml:"passwd"`
}

type gormCfg struct {
	// can add more
	SkipDefaultTransaction bool `yaml:"skip_default_transaction"`
	QueryFields            bool `yaml:"query_fields"`
}

type chatgpt struct {
	Token     string `yaml:"token"`
	Model     string `yaml:"model"`
	MaxTokens int    `yaml:"max_tokens"`
	Role      string `yaml:"role"`
	Stream    bool   `yaml:"stream"`
	Timeout   int    `yaml:"timeout"`
}

// Init 读取配置文件
func Init(configFilePath string) {
	once.Do(func() {
		// 读配置文件
		configData, err := os.ReadFile(configFilePath)
		if err != nil {
			pwd, _ := os.Getwd()
			log.Fatalf("load config file error %v, pwd: %s.", err, pwd)
		}
		config = &Config{}

		injectedConfigData := os.ExpandEnv(string(configData))

		config = &Config{}
		if err = yaml.Unmarshal([]byte(injectedConfigData), config); err != nil {
			log.Fatalf("unmarshal config file error %v.", err)
		}
	})
}

// Get 提供只读的全局配置
func Get() Config {
	return *config
}
