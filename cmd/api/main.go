package main

import (
	"fmt"
	"net/http"

	"github.com/AimrayX/weatherapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	if err := handlers.InitDB("/home/weatherstation/code/weather.db"); err != nil {
        log.Fatal(err)
        }

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API...")

	err := http.ListenAndServe("0.0.0.0:8000", r)
	if err != nil {
		log.Error(err)
	}
}
