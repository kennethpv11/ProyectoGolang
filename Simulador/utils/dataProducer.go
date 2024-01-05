package utils

import (
	"math/rand"
	"simulador/models"

	"github.com/goccy/go-json"
	"github.com/gofrs/uuid"
)

func GenerateDataElectricity() ([]byte, error) {
	id, err := uuid.NewV1()
	if err != nil {
		return nil, err
	}
	randInt := rand.Intn(11) + 90
	randIntMax := rand.Intn(11) + 110
	payload := models.ElectronicDevice{
		ValueMin: randInt,
		ValueMax: randIntMax,
		Measure:  "kwh",
		Id:       id,
	}
	return json.Marshal(payload)

}
