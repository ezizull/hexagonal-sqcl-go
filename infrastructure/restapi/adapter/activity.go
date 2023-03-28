package adapter

import (
	"hexagonal-sqlc/infrastructure/repository/postgres/sqlc"
	activityController "hexagonal-sqlc/infrastructure/restapi/controllers/activity"
)

func ActivityAdapter(db *sqlc.Queries) *activityController.Controller {
	return &activityController.Controller{ActivityService: db}
}
