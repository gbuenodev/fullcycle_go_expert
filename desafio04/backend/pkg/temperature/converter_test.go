package temperature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"zero", 0, 32},
		{"positive", 25, 77},
		{"negative", -10, 14},
		{"boiling point", 100, 212},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CelsiusToFahrenheit(tt.celsius)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"absolute zero", -273.15, 0},
		{"zero", 0, 273.15},
		{"positive", 25, 298.15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CelsiusToKelvin(tt.celsius)
			assert.InDelta(t, tt.expected, result, 0.01)
		})
	}
}
