package http_food

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food/domain"
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
	foods, err := h.usecase.GetAllFoods(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, foods)
}

func (h *FoodHandler) SaveFood(c *gin.Context) {
	var req SaveFoodRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	foodId := uuid.New()
	params := domain.Food{
		ID:        foodId,
		ImageURL:  req.ImageURL,
		Comment:   req.Comment,
		Nutrients: req.Nutrients,
	}

	if err := h.usecase.SaveFood(c.Request.Context(), &params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, SaveFoodResponse{ID: foodId})
}
