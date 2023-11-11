package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	priority "github.com/marioTiara/todolistwebapi/Priority"
)

type priorityHandler struct {
	priorityService priority.Service
}

func NewPriorityHandler(priorityService priority.Service) *priorityHandler {
	return &priorityHandler{priorityService: priorityService}
}

func (h *priorityHandler) PostPriorityHandler(c *gin.Context) {
	var priorityRequest priority.AddPriorotyTaskRequest
	err := c.ShouldBindJSON(&priorityRequest)

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

	priority, err := h.priorityService.Create(priorityRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": priority,
	})
}
