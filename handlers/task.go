package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (h *taskHandler) PostTasksHandler(c *gin.Context) {
	var taskRequest task.AddTaskRequest
	err := c.ShouldBindJSON((&taskRequest))

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	task, err := h.taskService.Create(taskRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})

}

func (h *taskHandler) GetTaskByIDHandler(c *gin.Context) {
	StrID := c.Param("id")
	id, err := strconv.ParseUint(StrID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "invalid id",
		})
	}

	tasks := h.taskService.FindByID(uint(id))
	if len(tasks) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})

}
