package activity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	"skyshi-gethired.go/infrastructure/restapi/controllers"
)

type Controller struct {
	ActivityService *sqlc.Queries
}

// GetActivities function return all activities
func (c *Controller) GetActivities(ctx *gin.Context) {
	var (
		activityResp []sqlc.Activity
		err          error
	)

	activityResp, err = c.ActivityService.GetActivities(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	activities := arrayToDomainMapper(activityResp)
	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}
