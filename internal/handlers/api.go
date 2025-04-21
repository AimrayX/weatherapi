package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/AimrayX/weatherapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Use(middleware.WithCORS)
	r.Route("/weatherdata", func(router chi.Router) {

		router.Use(middleware.Authorization)
		router.Get("/data", GetWeatherData)
	})
}