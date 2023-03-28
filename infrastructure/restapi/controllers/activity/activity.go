package activity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	domainActivity "skyshi-gethired.go/domain/activity"
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

// GetSingleActivity function return error when activity not found
func (c *Controller) GetSingleActivity(ctx *gin.Context) {
	var (
		activityResp sqlc.Activity
		err          error
	)

	activityIDStr := ctx.Param("id")
	activityID, err := strconv.ParseInt(activityIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Activity with ID " + activityIDStr + " Not Found"),
		})
		return
	}

	activityResp, err = c.ActivityService.GetActivityByID(ctx, int64(activityID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Activity with ID " + activityIDStr + " Not Found"),
		})
		return
	}

	activity := toDomainMapper(activityResp)
	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

// CreateSingleActivity function create a activity based newactivity body
func (c *Controller) CreateSingleActivity(ctx *gin.Context) {
	var (
		activityBody domainActivity.NewActivity
		activityResp sqlc.Activity
		err          error
	)

	// validation create activity body
	activityBody, message := createValidation(ctx)
	if message != "" {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: (message + " cannot be null"),
		})
		return
	}

	activityResp, err = c.ActivityService.CreateActivity(ctx, fromNewDomainMapper(&activityBody, "very-high"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	activitys := toDomainMapper(activityResp)
	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activitys,
	})
}
