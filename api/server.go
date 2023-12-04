package api

import (
	"github.com/gin-gonic/gin"
	task "github.com/marioTiara/todolistwebapi/Task"
	"github.com/marioTiara/todolistwebapi/utils"
)

type Server struct {
	config  utils.Config
	repo    task.Repository
	service task.Service
	router  *gin.Engine
}

func NewServer(config utils.Config, repo task.Repository, service task.Service) (*Server, error) {
	server := &Server{
		config:  config,
		repo:    repo,
		service: service,
	}

	return server, nil
}
