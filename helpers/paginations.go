package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marioTiara/todolistwebapi/dtos"
)

func GeneratePaginationRequest(context *gin.Context) *dtos.Pagination {
	//Convert query parameter string to int

	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(context.DefaultQuery("paege", "1"))

	sort := context.DefaultQuery("sort", "created_at desc")

	return &dtos.Pagination{Limit: limit, Page: page, Sort: sort}
}
