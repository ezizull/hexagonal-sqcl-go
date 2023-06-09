package todo

import (
	domainTodo "hexagonal-sqlc/domain/todo"

	"github.com/gin-gonic/gin"
)

func createValidation(ctx *gin.Context) (todoBody domainTodo.NewTodo, message string) {
	// Get body data for newtodo
	_ = ctx.BindJSON(&todoBody)

	if todoBody.Title == nil {
		return todoBody, "title"
	}

	if todoBody.ActivityGroupID == nil || *todoBody.ActivityGroupID == 0 {
		return todoBody, "activity_group_id"
	}

	if todoBody.IsActive == nil {
		return todoBody, "is_active"
	}

	return todoBody, message
}
