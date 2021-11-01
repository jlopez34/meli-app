package rest

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest  = MissionError{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Cannot process current request"}
	ErrInvalidJSON = MissionError{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
)

type MissionError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (e MissionError) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
