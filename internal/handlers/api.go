package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/AimrayX/piweatherstation/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Use(middleware.WithCors)

	r.Route("/weatherdata", func(router chi.Router){
		router.Get("/", GetWeatherData())
	})
}