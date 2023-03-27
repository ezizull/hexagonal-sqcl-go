package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	"skyshi-gethired.go/infrastructure/restapi/controllers"
)

type Controller struct {
	TodoService *sqlc.Queries
}

func (c *Controller) GetTodos(ctx *gin.Context) {
	activityGroupIDStr := ctx.DefaultQuery("activity_group_id", "0")
	activityGroupID, err := strconv.ParseInt(activityGroupIDStr, 10, 32)
	if err != nil {
		// handle the error if the string cannot be parsed as an int32
		activityGroupID = 0
	}

	todoResp, err := c.TodoService.GetTodosByActivity(ctx, int32(activityGroupID))
	if err != nil {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	todos := arrayToDomainMapper(todoResp)
	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "success",
		Message: "success",
		Data:    todos,
	})
}
