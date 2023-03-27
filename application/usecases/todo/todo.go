package todo

import (
	todoRepository "skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
)

type Service struct {
	TodoRepository todoRepository.Repository
}
