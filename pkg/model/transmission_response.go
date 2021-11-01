package model

type TransmissionResponse struct {
	Position *Position `json:"position,omitempty"`

	Message string `json:"message,omitempty"`
}
