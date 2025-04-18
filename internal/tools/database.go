package tools

import (
	log "github.com/sirupsen/logrus"
)

type WeatherDetails struct {
	Temperature int
	Humidity int
}

type DatabaseInterface interface {
	GetWeatherData(day string) *WeatherDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &weatherDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}