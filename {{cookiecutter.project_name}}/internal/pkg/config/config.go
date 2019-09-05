package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DatabateConfig struct {
	Host     string `mapstructure:"host"`
	Name     string `mapstrucure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ServerConfig struct {
	Port      string `mapstructure:"port"`
	Benchmark bool   `mapstructure:"benchmark"`
}

type Config struct {
	Db     DatabateConfig `mapstructure:"database"`
	Server ServerConfig   `mapstrucure:"server"`
}

func LoadConfig() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Couldn't load config: %s", err)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("couldn't read config: %s", err)
	}
	c.loadEnvConfig()
	return &c
}

func (c *Config) loadEnvConfig() {
	if port, ok := os.LookupEnv("PORT"); ok {
		c.Server.Port = port
	}
	if bm, ok := os.LookupEnv("BENCHMARK"); ok {
		c.Server.Benchmark = bm == "true"
	}
}

func LoadTestConfig(path string) *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("couldn't load config: %s", err)
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("couldn't read config: %s", err)
	}
	c.loadEnvConfig()
	return &c
}
