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
	foods, err := h.usecase.GetAllFoods(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, GetAllFoodResponse{Foods: foods})
}

func (h *FoodHandler) AnalyzeAndSaveFood(c *gin.Context) {

	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot open image"})
		return
	}
	defer file.Close()

	comment := c.PostForm("comment")
	params := &usecase.AnalyzeAndSaveFoodParams{
		File:     file,
		Filename: "",
		Comment:  comment,
	}

	resp, err := h.usecase.AnalyzeAndSaveFood(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, AnalyzeAndSaveFoodResponse{ID: resp.ID})
}
