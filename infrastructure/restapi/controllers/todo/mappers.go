package todo

import (
	"database/sql"

	domainTodo "skyshi-gethired.go/domain/todo"
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
)

func toDomainMapper(todo sqlc.Todo) *domainTodo.Todo {
	return &domainTodo.Todo{
		ID:              todo.ID,
		Title:           todo.Title.String,
		ActivityGroupID: todo.ActivityGroupID.Int32,
		IsActive:        todo.IsActive.Bool,
		CreatedAt:       todo.CreatedAt.Time,
		UpdatedAt:       todo.UpdatedAt.Time,
	}
}

func fromDomainMapper(todo domainTodo.Todo) *sqlc.Todo {
	return &sqlc.Todo{
		ID:              todo.ID,
		ActivityGroupID: sql.NullInt32{Int32: int32(todo.ActivityGroupID), Valid: true},
		Title:           sql.NullString{String: todo.Title, Valid: true},
		IsActive:        sql.NullBool{Bool: todo.IsActive, Valid: true},
		Priority:        sql.NullString{String: todo.Priority, Valid: true},
		CreatedAt:       sql.NullTime{Time: todo.CreatedAt, Valid: true},
		UpdatedAt:       sql.NullTime{Time: todo.UpdatedAt, Valid: true},
	}
}

func arrayToDomainMapper(Todo []sqlc.Todo) []*domainTodo.Todo {
	var todos []*domainTodo.Todo
	for _, todoResp := range Todo {
		todo := toDomainMapper(todoResp)
		todos = append(todos, todo)
	}
	return todos
}
