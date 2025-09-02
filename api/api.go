package api

import (
	"encoding/json"
	"net/http"
	"time"

)

type WeatherDetailParams struct {
	Day string
}

type WeatherDataResponse struct {
	
	Code int
	
	Day string
	Time time.Time
	WindSpeed int
	WindDirection int
	BarometricPressure int
	VOC int
	VCS int
	PM1dot0 float64
	PM1dot5 float64
	PM10 float64
	UVIndex int
	VisibleLight int
	InfraredLight int
	Temperature int
	Humidity int
	Rain int
}

type Error struct {

	Code int

	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error has occurred", http.StatusInternalServerError)
	}
)