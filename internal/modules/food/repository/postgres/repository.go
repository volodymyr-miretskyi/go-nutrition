package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/domain"
)

type FoodRepository struct {
	db *pgxpool.Pool
}

func NewFoodRepository(db *pgxpool.Pool) *FoodRepository {
	return &FoodRepository{
		db: db,
	}
}

func (r *FoodRepository) GetAll(ctx context.Context) ([]domain.Food, error) {
	query := `
		SELECT *
		FROM foods
	`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var foods []domain.Food

	for rows.Next() {
		var food domain.Food

		err := rows.Scan(
			&food.ID,
			&food.ImageURL,
			&food.Comment,
			&food.Nutrients.Calories,
			&food.Nutrients.Proteins,
			&food.Nutrients.Fats,
			&food.Nutrients.Carbs,
			&food.CreatedAt,
			&food.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func (r *FoodRepository) Save(ctx context.Context, food *domain.Food) error {
	query := `
	INSERT INTO foods (
		id,
		image_url,
		comment,
		calories,
		proteins,
		fats,
		carbs
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	)
	`
	_, err := r.db.Exec(
		ctx, query,
		food.ID,
		food.ImageURL,
		food.Comment,
		food.Nutrients.Calories,
		food.Nutrients.Proteins,
		food.Nutrients.Fats,
		food.Nutrients.Carbs,
	)

	return err
}
