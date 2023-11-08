package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	task "github.com/marioTiara/todolistwebapi/Task"
)

type taskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *taskHandler {
	return &taskHandler{taskService}
}
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Mario Tiara",
		"bio":  "A Full Stack SOftware Enginner",
	})
}

func (handler *taskHandler) TasksHandler(c *gin.Context) {

}
