package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	DBName   string
	Host     string
	User     string
	Password string
	Port     string
	SSLMode  string
}

type LoggerConfig struct {
	File string
}

func NewConfig(path string) Config {
	var cfg Config

	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Errorf("reading config error: %w", err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println(fmt.Errorf("ummarshal to config struct is failed: %w", err))
	}

	return cfg
}
