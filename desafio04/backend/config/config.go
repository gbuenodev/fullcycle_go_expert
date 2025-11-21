package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig
	WeatherAPI WeatherAPIConfig
	ViaCEP     ViaCEPConfig
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type WeatherAPIConfig struct {
	Key     string `mapstructure:"key"`
	BaseURL string `mapstructure:"base_url"`
}

type ViaCEPConfig struct {
	BaseURL string `mapstructure:"base_url"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path + "/.env")
	viper.SetConfigType("env")

	viper.SetDefault("PORT", "3000")
	viper.SetDefault("VIACEP_BASE_URL", "https://viacep.com.br/ws")
	viper.SetDefault("WEATHER_API_BASE_URL", "https://api.weatherapi.com/v1")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading .env file: %w", err)
		}
	}

	// AutomaticEnv must be called AFTER ReadInConfig so environment variables take precedence
	viper.AutomaticEnv()

	cfg := &Config{
		Server: ServerConfig{
			Port: viper.GetString("PORT"),
		},
		WeatherAPI: WeatherAPIConfig{
			Key:     viper.GetString("WEATHER_API_KEY"),
			BaseURL: viper.GetString("WEATHER_API_BASE_URL"),
		},
		ViaCEP: ViaCEPConfig{
			BaseURL: viper.GetString("VIACEP_BASE_URL"),
		},
	}

	if cfg.WeatherAPI.Key == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY is required")
	}

	return cfg, nil
}
