package rest

import (
	"encoding/json"
	"net/http"

	"github.com/jlopez34/meli-app/pkg/finding"
)

func handleFinding(as finding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to our handleFinding!")
	}
}

func handleFindingSingle(as finding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to our handleFindingSingle!")
	}
}
