package adapter

import (
	todoService "skyshi-gethired.go/application/usecases/todo"
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	todoRepository "skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	todoController "skyshi-gethired.go/infrastructure/restapi/controllers/todo"
)

func TodoAdapter(db *sqlc.Queries) *todoController.Controller {
	mRepository := todoRepository.Repository{DB: db}
	service := todoService.Service{TodoRepository: mRepository}
	return &todoController.Controller{BookService: service}
}
