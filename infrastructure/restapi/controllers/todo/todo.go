package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	"skyshi-gethired.go/infrastructure/restapi/controllers"
	"skyshi-gethired.go/utils"
)

type Controller struct {
	TodoService *sqlc.Queries
}

// GetTodos function return all todos when activity_group_id = 0 or empty
func (c *Controller) GetTodos(ctx *gin.Context) {
	var (
		todoResp []sqlc.Todo
		err      error
	)

	activityGroupIDStr := ctx.DefaultQuery("activity_group_id", "0")
	if activityGroupIDStr != "0" {
		activityGroupID := utils.ConvertToNullInt32(activityGroupIDStr)
		todoResp, err = c.TodoService.GetTodosByActivity(ctx, activityGroupID)
		if err != nil {
			ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}
	}

	todoResp, err = c.TodoService.GetAllTodos(ctx)
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

// GetSingleTodo function return error when todos not found
func (c *Controller) GetSingleTodo(ctx *gin.Context) {
	var (
		todoResp sqlc.Todo
		err      error
	)

	todoIDStr := ctx.Param("id")
	todoID, err := strconv.ParseInt(todoIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoIDStr + " Not Found"),
		})
		return
	}

}
