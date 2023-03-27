package todo

import (
	"database/sql"

	domainTodo "skyshi-gethired.go/domain/todo"
)

func (todo *Todo) ToDomainMapper() *domainTodo.Todo {
	return &domainTodo.Todo{
		ID:              todo.ID,
		Title:           todo.Title.String,
		ActivityGroupID: todo.ActivityGroupID.Int32,
		IsActive:        todo.IsActive.Bool,
		CreatedAt:       todo.CreatedAt.Time,
		UpdatedAt:       todo.UpdatedAt.Time,
	}
}

func FromDomainMapper(todo *domainTodo.Todo) *Todo {
	return &Todo{
		ID:              todo.ID,
		ActivityGroupID: sql.NullInt32{Int32: int32(todo.ActivityGroupID), Valid: true},
		Title:           sql.NullString{String: todo.Title, Valid: true},
		IsActive:        sql.NullBool{Bool: todo.IsActive, Valid: true},
		Priority:        sql.NullString{String: todo.Priority, Valid: true},
		CreatedAt:       sql.NullTime{Time: todo.CreatedAt, Valid: true},
		UpdatedAt:       sql.NullTime{Time: todo.UpdatedAt, Valid: true},
	}
}

func ArrayToDomainMapper(todos *[]Todo) *[]domainTodo.Todo {
	todosDomain := make([]domainTodo.Todo, len(*todos))
	for i, todo := range *todos {
		todosDomain[i] = *todo.ToDomainMapper()
	}

	return &todosDomain
}
