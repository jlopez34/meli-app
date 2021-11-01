package rest

import (
	"encoding/json"
	"net/http"
)

type ResponseAPI struct {
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func Success(result interface{}, status int) *ResponseAPI {
	return &ResponseAPI{
		Success: true,
		Status:  status,
		Result:  result,
	}
}

func (r *ResponseAPI) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}
