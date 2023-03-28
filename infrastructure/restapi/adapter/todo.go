package adapter

import (
	"hexagonal-sqlc/infrastructure/repository/postgres/sqlc"
	todoController "hexagonal-sqlc/infrastructure/restapi/controllers/todo"
)

func TodoAdapter(db *sqlc.Queries) *todoController.Controller {
	return &todoController.Controller{TodoService: db}
}
