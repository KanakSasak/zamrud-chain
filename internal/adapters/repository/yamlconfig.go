package repository

import (
	"github.com/spf13/viper"
	"zamrud/internal/core/ports"
)

type ConfigRepository struct{}

func NewConfigRepository() ports.ConfigRepository {
	return &ConfigRepository{}
}

func (c *ConfigRepository) LoadConfig() (ports.Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	//viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return ports.Config{}, err
	}

	var config ports.Config
	if err := viper.Unmarshal(&config); err != nil {
		return ports.Config{}, err
	}

	return config, nil
}
