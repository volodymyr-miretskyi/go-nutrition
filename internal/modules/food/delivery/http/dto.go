package http_food

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/domain"
)

type GetAllFoodResponse struct {
	Foods []domain.Food `json:"foods"`
}

type AnalyzeAndSaveFoodRequest struct {
	Image   multipart.File `json:"image"`
	Comment string         `json:"comment"`
}

type AnalyzeAndSaveFoodResponse struct {
	ID uuid.UUID `json:"id"`
}
