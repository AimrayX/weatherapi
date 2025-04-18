package middleware

import (
	"errors"
	"net/http"

	"github.com/AimrayX/weatherapi/api"
	"github.com/AimrayX/weatherapi/internal/tools"
	log "github.com/sirupsen/logrus"

)

var UnauthorizedError = errors.New("Invalid API token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		var token = r.Header.Get("Authorization")
		var err error


		if token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var weatherDetails *tools.weatherDetails
		weatherDetails = (*database).GetWeatherData()

		if (weatherDetails == nil || (token != )) {
			
		}
	})
}