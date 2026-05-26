package http_food

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/usecase"
)

type FoodHandler struct {
	usecase *usecase.FoodUsecase
}

func NewFoodHandler(u *usecase.FoodUsecase) *FoodHandler {
	return &FoodHandler{
		usecase: u,
	}
}

func (h *FoodHandler) GetFoods(c *gin.Context) {
	foods, err := h.usecase.GetAllFoods(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, foods)
}

func (h *FoodHandler) AnalyzeFood() {

}
