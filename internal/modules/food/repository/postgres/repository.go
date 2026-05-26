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
			&food.Nutrients,
		)

		if err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}
