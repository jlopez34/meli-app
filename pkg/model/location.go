package model

type Location struct {
	Distance float32 `json:"distance,omitempty"`

	Message []string `json:"message,omitempty"`
}
