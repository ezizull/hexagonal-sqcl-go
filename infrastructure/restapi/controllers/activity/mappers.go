package activity

import (
	"database/sql"

	domainActivity "hexagonal-sqlc/domain/activity"
	"hexagonal-sqlc/infrastructure/repository/postgres/sqlc"
)

func fromUpdateDomainMapper(todo *domainActivity.UpdateActivity, id int64) sqlc.UpdateActivityParams {
	return sqlc.UpdateActivityParams{
		ID:    id,
		Title: sql.NullString{String: *todo.Title, Valid: true},
	}
}

func fromNewDomainMapper(todo *domainActivity.NewActivity, todoPriority string) sqlc.CreateActivityParams {
	return sqlc.CreateActivityParams{
		Title: sql.NullString{String: *todo.Title, Valid: true},
		Email: sql.NullString{String: *todo.Email, Valid: true},
	}
}

func toDomainMapper(todo sqlc.Activity) *domainActivity.Activity {
	return &domainActivity.Activity{
		ID:        todo.ID,
		Title:     todo.Title.String,
		Email:     todo.Email.String,
		CreatedAt: todo.CreatedAt.Time,
		UpdatedAt: todo.UpdatedAt.Time,
	}
}

func fromDomainMapper(todo domainActivity.Activity) *sqlc.Activity {
	return &sqlc.Activity{
		ID:        todo.ID,
		Title:     sql.NullString{String: todo.Title, Valid: true},
		Email:     sql.NullString{String: todo.Email, Valid: true},
		CreatedAt: sql.NullTime{Time: todo.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: todo.UpdatedAt, Valid: true},
	}
}

func arrayToDomainMapper(Todo []sqlc.Activity) []*domainActivity.Activity {
	var todos []*domainActivity.Activity
	for _, todoResp := range Todo {
		todo := toDomainMapper(todoResp)
		todos = append(todos, todo)
	}
	return todos
}
