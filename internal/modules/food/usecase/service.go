package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/domain"
)

type FoodUsecase struct {
	repo     FoodRepository
	analyzer FoodAnalyzer
	storage  ImageStorage
}

func NewFoodUsecase(r FoodRepository, a FoodAnalyzer, s ImageStorage) *FoodUsecase {
	return &FoodUsecase{
		repo:     r,
		analyzer: a,
		storage:  s,
	}
}

func (u *FoodUsecase) GetAllFoods(ctx context.Context) ([]domain.Food, error) {
	return u.repo.GetAll(ctx)
}

func (u *FoodUsecase) AnalyzeAndSaveFood(ctx context.Context, params *AnalyzeAndSaveFoodParams) (*domain.Food, error) {

	// Save image to S3

	analyzeFoodParams := FoodAnalyzerAnalyzeFoodParams{
		ImageUrl: "https://img.magnific.com/free-psd/roasted-chicken-dinner-platter-delicious-feast_632498-25445.jpg?semt=ais_hybrid&w=740&q=80", // replace
		Comment:  params.Comment,
	}

	nutrients, err := u.analyzer.AnalyzeFood(ctx, &analyzeFoodParams)

	if err != nil {
		return nil, err
	}

	foodId := uuid.New()
	food := domain.Food{
		ID:        foodId,
		ImageURL:  "https://img.magnific.com/free-psd/roasted-chicken-dinner-platter-delicious-feast_632498-25445.jpg?semt=ais_hybrid&w=740&q=80", // Replace
		Comment:   params.Comment,
		Nutrients: *nutrients,
	}

	if err := u.repo.Save(ctx, &food); err != nil {
		return nil, err
	}

	return &food, nil
}
