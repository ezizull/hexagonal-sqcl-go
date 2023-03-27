package adapter

import (
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	todoController "skyshi-gethired.go/infrastructure/restapi/controllers/todo"
)

func TodoAdapter(db *sqlc.Queries) *todoController.Controller {
	return &todoController.Controller{TodoService: db}
}
