package config

import (
	"github.com/spf13/viper"
)

type NetServerConfig struct {
	Host string
	Port int
}

type RGeneratorConfig struct {
	RoutePerSec int
	WorkTimeSec int
}

type Config struct {
	ServerListenerTcp NetServerConfig
	RGenerator        RGeneratorConfig
}

func parseConfigFile(configDir string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("serverListener.tcp", &cfg.ServerListenerTcp); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("requestGenerator", &cfg.RGenerator); err != nil {
		return err
	}
	return nil
}

func Init(configDir string) (*Config, error) {
	viper.SetConfigName("config")
	if err := parseConfigFile(configDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
