package http_food

import (
	"github.com/gin-gonic/gin"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/usecase"
)

func RegisterFoodRoutes(router *gin.RouterGroup, u *usecase.FoodUsecase) {
	handler := NewFoodHandler(u)

	food := router.Group("/foods")

	{
		food.GET("/", handler.GetFoods)
	}
}
