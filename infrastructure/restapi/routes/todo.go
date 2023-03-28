package routes

import (
	"github.com/gin-gonic/gin"
	todoController "skyshi-gethired.go/infrastructure/restapi/controllers/todo"
)

func TodoRoutes(router *gin.Engine, controller *todoController.Controller) {
	routerTodo := router.Group("/todo-items")
	{
		routerTodo.GET("", controller.GetTodos)
		routerTodo.GET("/:id", controller.GetSingleTodo)
		routerTodo.POST("", controller.CreateSingleTodo)
		routerTodo.PATCH("/:id", controller.UpdateSingleTodo)
	}

}
