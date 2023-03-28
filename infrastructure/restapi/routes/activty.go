package routes

import (
	activityController "hexagonal-sqlc/infrastructure/restapi/controllers/activity"

	"github.com/gin-gonic/gin"
)

func ActivityRoutes(router *gin.Engine, controller *activityController.Controller) {
	routerTodo := router.Group("/activity-groups")
	{
		routerTodo.GET("", controller.GetActivities)
		routerTodo.GET("/:id", controller.GetSingleActivity)
		routerTodo.POST("", controller.CreateSingleActivity)
		routerTodo.PATCH("/:id", controller.UpdateSingleActivity)
		routerTodo.DELETE("/:id", controller.DeleteSingleActivity)
	}

}
