package rest

import (
	"github.com/go-chi/chi"
	"github.com/jlopez34/meli-app/pkg/finding"
)

func Handler(as finding.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/topsecret", handleFinding(as))
	router.Post("/topsecret_split/:satellite_name", handleFindingSingle(as))

	return router
}
