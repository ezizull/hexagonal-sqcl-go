package routes

import (
	"github.com/gin-gonic/gin"
	activityController "skyshi-gethired.go/infrastructure/restapi/controllers/activity"
)

func ActivityRoutes(router *gin.Engine, controller *activityController.Controller) {
	routerTodo := router.Group("/activity-groups")
	{
		routerTodo.GET("", controller.GetActivities)
	}

}