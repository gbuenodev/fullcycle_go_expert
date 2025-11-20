package entity

type Address struct {
	City    string
	Weather *Weather
}

func NewAddress(city string) *Address {
	return &Address{
		City:    city,
		Weather: nil,
	}
}

func (a *Address) SetWeather(tempCelsius float64) {
	a.Weather = NewWeather(tempCelsius)
}
