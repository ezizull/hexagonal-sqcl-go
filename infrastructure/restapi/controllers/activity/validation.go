package activity

import (
	domainActivity "hexagonal-sqlc/domain/activity"

	"github.com/gin-gonic/gin"
)

func createValidation(ctx *gin.Context) (activityBody domainActivity.NewActivity, message string) {
	// Get body data for newtodo
	_ = ctx.BindJSON(&activityBody)

	if activityBody.Title == nil {
		return activityBody, "title"
	}

	if activityBody.Email == nil {
		return activityBody, "email"
	}

	return activityBody, message
}

func updateValidation(ctx *gin.Context) (activityBody domainActivity.UpdateActivity, message string) {
	// Get body data for newtodo
	_ = ctx.BindJSON(&activityBody)

	if activityBody.Title == nil {
		return activityBody, "title"
	}

	return activityBody, message
}
