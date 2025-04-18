package tools

import (
	"time"
)

type weatherDB struct {}

var mockWeatherDB = map[string]WeatherDetails{
	"today": {
		Temperature: 10,
		Humidity: 53,
	},
	"yesterday": {
		Temperature: 19,
		Humidity: 42,
	},
}

func (d *weatherDB) GetWeatherData(day string) *WeatherDetails {

	time.Sleep(1000)

	var weatherData = WeatherDetails{}
	weatherData, ok := mockWeatherDB[day]
	if !ok {
		return nil
	}

	return &weatherData
}

func (d *weatherDB) SetupDatabase() error {
	return nil
}