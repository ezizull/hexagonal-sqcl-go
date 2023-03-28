package routes

import (
	todoController "hexagonal-sqlc/infrastructure/restapi/controllers/todo"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine, controller *todoController.Controller) {
	routerTodo := router.Group("/todo-items")
	{
		routerTodo.GET("", controller.GetTodos)
		routerTodo.GET("/:id", controller.GetSingleTodo)
		routerTodo.POST("", controller.CreateSingleTodo)
		routerTodo.PATCH("/:id", controller.UpdateSingleTodo)
		routerTodo.DELETE("/:id", controller.DeleteSingleTodo)
	}

}
