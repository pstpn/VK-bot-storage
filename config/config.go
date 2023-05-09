package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	configPath             = "config/config.yaml"
	DeletePasswordDeadline = 20 * time.Second
)

type Config struct {
	Bot      BotConfig      `yaml:"bot"`
	Database DatabaseConfig `yaml:"database"`
}

type BotConfig struct {
	Token string `yaml:"token"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewConfig() (*Config, error) {

	var err error
	var config Config

	viper.SetConfigFile(configPath)

	err = viper.BindEnv("database.host", "DB_HOST")
	if err != nil {
		return nil, err
	}
	err = viper.BindEnv("database.port", "DB_PORT")
	if err != nil {
		return nil, err
	}
	err = viper.BindEnv("database.user", "DB_USER")
	if err != nil {
		return nil, err
	}
	err = viper.BindEnv("database.password", "DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	err = viper.BindEnv("database.database", "DB_DATABASE")
	if err != nil {
		return nil, err
	}
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
