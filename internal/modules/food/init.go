package food

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	http_food "github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/delivery/http"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/repository/postgres"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/usecase"
)

func InitModule(router *gin.RouterGroup, db *pgxpool.Pool) {
	repo := postgres.NewFoodRepository(db)
	uc := usecase.NewFoodUsecase(repo)

	http_food.RegisterFoodRoutes(router, uc)
}
