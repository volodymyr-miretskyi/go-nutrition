package domain

import (
	"github.com/google/uuid"
)

type Food struct {
	ID        uuid.UUID
	ImageURL  string
	Comment   string
	Nutrients Nutrients
}

type Nutrients struct {
	Calories string
	Protein  string
	Carbs    string
	Fat      string
}
