package todo

import (
	"net/http"
	"strconv"

	domainTodo "skyshi-gethired.go/domain/todo"

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
				Status:  "Error",
				Message: err.Error(),
			})
			return
		}
	}

	todoResp, err = c.TodoService.GetAllTodos(ctx)
	if err != nil {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	todos := arrayToDomainMapper(todoResp)
	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
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
	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoIDStr + " Not Found"),
		})
		return
	}

	todoResp, err = c.TodoService.GetTodosByID(ctx, int64(todoID))
	if err != nil {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoIDStr + " Not Found"),
		})
		return
	}

	todos := toDomainMapper(todoResp)
	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

// CreateSingleTodo function create a todo based newtodo body
func (c *Controller) CreateSingleTodo(ctx *gin.Context) {
	var (
		todoBody domainTodo.NewTodo
		todoResp sqlc.Todo
		err      error
	)

	// Get body data for newtodo
	_ = ctx.BindJSON(&todoBody)

	if todoBody.Title == nil {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: ("title cannot be null"),
		})
		return
	}

	if todoBody.ActivityGroupID == nil || *todoBody.ActivityGroupID == 0 {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: ("activity_group_id cannot be null"),
		})
		return
	}

	if todoBody.IsActive == nil {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: ("is_active cannot be null"),
		})
		return
	}

	todoResp, err = c.TodoService.CreateTodo(ctx, fromNewDomainMapper(&todoBody, "very-high"))
	if todoResp.ID == 0 {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Activity Group with ID " + strconv.Itoa(int(*todoBody.ActivityGroupID)) + " Not Found"),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusAccepted, controllers.ErrorResponse{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	todos := toDomainMapper(todoResp)
	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

// UpdateTodo function update a todo based updatetodo body
