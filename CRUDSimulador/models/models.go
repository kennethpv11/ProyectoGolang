// Package models contains all the models in the program
package models

import (
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
)

// ElectricityInfo represent the model in the database
type ElectricityInfo struct {
	bun.BaseModel `bun:"table:col_electricity_info.metric,alias:u"`
	ValueMin      int       `bun:"value_min"`
	ValueMax      int       `bun:"value_max"`
	Measure       string    `bun:"measure"`
	ID            uuid.UUID `bun:"id,pk"`
}
