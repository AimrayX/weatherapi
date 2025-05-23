package middleware

import (
	"errors"
	"net/http"

	"github.com/AimrayX/weatherapi/api"
	//"github.com/AimrayX/weatherapi/internal/tools"
	log "github.com/sirupsen/logrus"

)

var ErrUnauthorizedError = errors.New("invalid API token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		var token = r.Header.Get("Authorization")

		if token != "1234567890" {
			log.Error(ErrUnauthorizedError)
			api.RequestErrorHandler(w, ErrUnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
		
	})
}