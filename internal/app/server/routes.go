package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/modules/food"
)

func (s *Server) mountRoutes() {
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := s.router.Group("/api/v1")

	food.InitModule(api, s.db)
}
