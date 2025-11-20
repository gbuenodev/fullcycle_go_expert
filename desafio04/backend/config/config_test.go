package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func resetViper() {
	viper.Reset()
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name        string
		envContent  string
		expectError bool
		errorMsg    string
		validate    func(t *testing.T, cfg *Config)
	}{
		{
			name: "valid env file",
			envContent: `PORT=8080
WEATHER_API_KEY=test-key-123
WEATHER_API_BASE_URL=https://test.weatherapi.com
VIACEP_BASE_URL=https://test.viacep.com
`,
			expectError: false,
			validate: func(t *testing.T, cfg *Config) {
				assert.Equal(t, "8080", cfg.Server.Port)
				assert.Equal(t, "test-key-123", cfg.WeatherAPI.Key)
				assert.Equal(t, "https://test.weatherapi.com", cfg.WeatherAPI.BaseURL)
				assert.Equal(t, "https://test.viacep.com", cfg.ViaCEP.BaseURL)
			},
		},
		{
			name:        "empty env file",
			envContent:  "",
			expectError: true,
			errorMsg:    "WEATHER_API_KEY is required",
		},
		{
			name: "env file with empty WEATHER_API_KEY",
			envContent: `PORT=3000
WEATHER_API_KEY=
WEATHER_API_BASE_URL=https://api.weatherapi.com/v1
VIACEP_BASE_URL=https://viacep.com.br/ws
`,
			expectError: true,
			errorMsg:    "WEATHER_API_KEY is required",
		},
		{
			name:        "env file with only WEATHER_API_KEY loads defaults",
			envContent:  `WEATHER_API_KEY=my-api-key`,
			expectError: false,
			validate: func(t *testing.T, cfg *Config) {
				assert.Equal(t, "my-api-key", cfg.WeatherAPI.Key)
				assert.Equal(t, "3000", cfg.Server.Port)
				assert.Equal(t, "https://api.weatherapi.com/v1", cfg.WeatherAPI.BaseURL)
				assert.Equal(t, "https://viacep.com.br/ws", cfg.ViaCEP.BaseURL)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetViper()

			tmpDir := t.TempDir()
			err := os.WriteFile(tmpDir+"/.env", []byte(tt.envContent), 0644)
			assert.NoError(t, err, "Failed to create .env file")

			cfg, err := LoadConfig(tmpDir)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
				assert.Nil(t, cfg)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, cfg)
				if tt.validate != nil {
					tt.validate(t, cfg)
				}
			}
		})
	}
}
