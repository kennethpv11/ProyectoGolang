package models

import (
	"github.com/gofrs/uuid"
)

type ElectronicDevice struct {
	ValueMin int       `json:"value_min"`
	ValueMax int       `json:"value_max"`
	Measure  string    `json:"measure"`
	Id       uuid.UUID `json:"id"`
}

type ElectricityResponse struct {
	Message string
}
