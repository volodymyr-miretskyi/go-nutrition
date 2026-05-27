package usecase

import (
	"context"

	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/domain"
)

type FoodUsecase struct {
	repo FoodRepository
}

func NewFoodUsecase(r FoodRepository) *FoodUsecase {
	return &FoodUsecase{
		repo: r,
	}
}

func (u *FoodUsecase) GetAllFoods(ctx context.Context) ([]domain.Food, error) {
	return u.repo.GetAll(ctx)
}

func (u *FoodUsecase) SaveFood(ctx context.Context, food *domain.Food) error {
	return u.repo.Save(ctx, food)
}
