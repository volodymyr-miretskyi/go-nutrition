package http_food

import (
	"github.com/google/uuid"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/domain"
)

type SaveFoodRequest struct {
	ImageURL  string           `json:"imageUrl"`
	Comment   string           `json:"comment"`
	Nutrients domain.Nutrients `json:"nutrients"`
}

type SaveFoodResponse struct {
	ID uuid.UUID
}
