package activity

import (
	activityRepository "skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
)

type Service struct {
	activityRepository activityRepository.Repository
}
