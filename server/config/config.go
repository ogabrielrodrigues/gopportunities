package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type serverConfig struct {
	Port     string
	Origin   string
	Database string
}

var config *serverConfig

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("config initialization error: %v", err)
	}

	config = &serverConfig{
		Port:     viper.GetString("server.port"),
		Origin:   viper.GetString("server.origin"),
		Database: viper.GetString("server.database"),
	}

	return nil
}

func GetServerConfig() *serverConfig {
	return config
}
