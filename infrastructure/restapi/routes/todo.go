package routes

import (
	"github.com/gin-gonic/gin"
	todoController "skyshi-gethired.go/infrastructure/restapi/controllers/todo"
)

func TodoRoutes(router *gin.RouterGroup, controller *todoController.Controller) {
	routerTodo := router.Group("/todo-items")
	{
		routerTodo.POST("", controller.GetTodos)
	}

}
