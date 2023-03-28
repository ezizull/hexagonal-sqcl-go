package adapter

import (
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	activityController "skyshi-gethired.go/infrastructure/restapi/controllers/activity"
)

func ActivityAdapter(db *sqlc.Queries) *activityController.Controller {
	return &activityController.Controller{ActivityService: db}
}
