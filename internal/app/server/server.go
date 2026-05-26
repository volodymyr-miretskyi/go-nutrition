package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/app/config"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
	db     *pgxpool.Pool
}

func New(cfg *config.Config, pool *pgxpool.Pool) *Server {
	router := gin.Default()

	s := &Server{
		cfg:    cfg,
		router: router,
		db:     pool,
	}

	s.mountRoutes()

	return s
}

func (s *Server) Run() error {
	runUrl := ":" + s.cfg.Port
	return s.router.Run(runUrl)
}
