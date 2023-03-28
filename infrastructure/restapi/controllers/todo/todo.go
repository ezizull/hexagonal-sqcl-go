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
			ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
				Status:  "Error",
				Message: err.Error(),
			})
			return
		}
	}

	todoResp, err = c.TodoService.GetAllTodos(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
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
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoIDStr + " Not Found"),
		})
		return
	}

	todoResp, err = c.TodoService.GetTodosByID(ctx, int64(todoID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
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

	// validation create todo body
	todoBody, message := createValidation(ctx)
	if message != "" {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: (message + " cannot be null"),
		})
		return
	}

	todoResp, err = c.TodoService.CreateTodo(ctx, fromNewDomainMapper(&todoBody, "very-high"))
	if todoResp.ID == 0 {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Activity Group with ID " + strconv.Itoa(int(*todoBody.ActivityGroupID)) + " Not Found"),
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
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

// UpdateSingleTodo function update a single todo based updatetodo body
func (c *Controller) UpdateSingleTodo(ctx *gin.Context) {
	var (
		todoBody domainTodo.UpdateTodo
		todoResp sqlc.Todo
	)

	todoIDStr := ctx.Param("id")
	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoIDStr + " Not Found"),
		})
		return
	}

	// Get body data for updatetodo
	err = ctx.BindJSON(&todoBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	// Get single todo for
	todoResp, err = c.TodoService.UpdateTodo(ctx, fromUpdateDomainMapper(&todoBody, int64(todoID)))
	if todoResp.ID == 0 {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + strconv.Itoa(int(todoID)) + " Not Found"),
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
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
