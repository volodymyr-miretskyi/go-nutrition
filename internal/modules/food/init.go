package food

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	analyzer "github.com/volodymyr-miretskyi/go-nutrition/internal/adapters/openai"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/app/config"
	http_food "github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/delivery/http"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/repository/postgres"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/usecase"
)

func InitModule(router *gin.RouterGroup, db *pgxpool.Pool, cfg *config.Config) {
	repo := postgres.NewFoodRepository(db)

	foodAnalyzer := analyzer.New(cfg.OpenAi.ApiKey)

	uc := usecase.NewFoodUsecase(repo, foodAnalyzer, nil) // Add S3 instead of nil

	http_food.RegisterFoodRoutes(router, uc)
}
