package entity

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

type Weather struct {
	Temperature
}

func NewWeather(tempCelsius float64) *Weather {
	return &Weather{
		Temperature: Temperature{
			Celsius:    tempCelsius,
			Fahrenheit: celsiusToFarenheit(tempCelsius),
			Kelvin:     celsiusToKelvin(tempCelsius),
		},
	}
}

func celsiusToFarenheit(tempCelsius float64) float64 {
	return tempCelsius*1.8 + 32
}

func celsiusToKelvin(tempCelsius float64) float64 {
	return tempCelsius + 273
}
