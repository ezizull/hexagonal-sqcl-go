package activity

import (
	"github.com/gin-gonic/gin"
	domainActivity "skyshi-gethired.go/domain/activity"
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
