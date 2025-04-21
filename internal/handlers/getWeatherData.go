package handlers

import (
	"html/template"
	"net/http"

	"github.com/AimrayX/weatherapi/api"
	"github.com/AimrayX/weatherapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
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

	// Assume we want to return HTML (you can check for r.URL.Query().Get("format") if needed)
	tmpl, err := template.ParseFiles("internal/templates/weather.html")
	if err != nil {
		log.Error("Template parsing error:", err)
		api.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, weatherDetails)
	if err != nil {
		log.Error("Template execution error:", err)
		api.InternalErrorHandler(w)
		return
	}
}