package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AimrayX/weatherapi/api"
	"github.com/AimrayX/weatherapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetWeatherData(w http.ResponseWriter, r *http.Request) {
	var params = api.WeatherParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var weatherDetails *tools.WeatherDetails
	weatherDetails = (*database).GetWeatherData(params.Day)

	if weatherDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.WeatherDataResponse{
		Temperature: (*weatherDetails).Temperature,
		Humidity: (*weatherDetails).Humidity,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}