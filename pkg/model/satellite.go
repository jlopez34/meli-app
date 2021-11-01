package model

type Satellite struct {
	Name string `json:"name,omitempty"`

	Distance float32 `json:"distance,omitempty"`

	Message []string `json:"message,omitempty"`
}
