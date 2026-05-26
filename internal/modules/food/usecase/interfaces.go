package usecase

import (
	"context"

	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/domain"
)

type FoodRepository interface {
	GetAll(ctx context.Context) ([]domain.Food, error)
}
