package domain

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	ID        uuid.UUID
	ImageURL  string
	Comment   string
	Nutrients Nutrients
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Nutrients struct {
	Calories int
	Proteins int
	Carbs    int
	Fats     int
}
