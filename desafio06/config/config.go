package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	RedisAddr                   string `mapstructure:"REDIS_ADDR"`
	RateLimitIP                 int    `mapstructure:"RATE_LIMIT_IP"`
	RateLimitIPBlockDuration    int    `mapstructure:"RATE_LIMIT_IP_BLOCK_DURATION"`
	RateLimitTokenBasic         int    `mapstructure:"RATE_LIMIT_TOKEN_BASIC"`
	RateLimitTokenPremium       int    `mapstructure:"RATE_LIMIT_TOKEN_PREMIUM"`
	RateLimitTokenBlockDuration int    `mapstructure:"RATE_LIMIT_TOKEN_BLOCK_DURATION"`
	RateLimitWindow             int    `mapstructure:"RATE_LIMIT_WINDOW"`
	ServerPort                  string `mapstructure:"SERVER_PORT"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	viper.BindEnv("REDIS_ADDR")
	viper.BindEnv("RATE_LIMIT_IP")
	viper.BindEnv("RATE_LIMIT_IP_BLOCK_DURATION")
	viper.BindEnv("RATE_LIMIT_TOKEN_BASIC")
	viper.BindEnv("RATE_LIMIT_TOKEN_PREMIUM")
	viper.BindEnv("RATE_LIMIT_TOKEN_BLOCK_DURATION")
	viper.BindEnv("RATE_LIMIT_WINDOW")
	viper.BindEnv("SERVER_PORT")

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if config.RedisAddr == "" {
		return nil, fmt.Errorf("REDIS_ADDR is required")
	}

	if config.RateLimitWindow <= 0 {
		config.RateLimitWindow = 1
	}

	return &config, nil
}
