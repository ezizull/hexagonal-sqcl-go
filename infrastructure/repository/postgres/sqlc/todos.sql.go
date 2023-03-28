// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: todos.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (activity_group_id, title, is_active, priority)
VALUES ($1, $2, $3, $4)
RETURNING id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at
`

type CreateTodoParams struct {
	ActivityGroupID sql.NullInt32
	Title           sql.NullString
	IsActive        sql.NullBool
	Priority        sql.NullString
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.ActivityGroupID,
		arg.Title,
		arg.IsActive,
		arg.Priority,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.ActivityGroupID,
		&i.Title,
		&i.IsActive,
		&i.Priority,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getAllTodos = `-- name: GetAllTodos :many
SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at FROM todos
ORDER BY id DESC
`

func (q *Queries) GetAllTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getAllTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.ActivityGroupID,
			&i.Title,
			&i.IsActive,
			&i.Priority,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodosByActivity = `-- name: GetTodosByActivity :many
SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at FROM todos
WHERE activity_group_id = $1
ORDER BY id DESC
`

func (q *Queries) GetTodosByActivity(ctx context.Context, activityGroupID sql.NullInt32) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getTodosByActivity, activityGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.ActivityGroupID,
			&i.Title,
			&i.IsActive,
			&i.Priority,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodosByID = `-- name: GetTodosByID :one
SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at FROM todos
WHERE id = $1
`

func (q *Queries) GetTodosByID(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodosByID, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.ActivityGroupID,
		&i.Title,
		&i.IsActive,
		&i.Priority,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :exec
INSERT INTO todos (activity_group_id, title, is_active)
VALUES ($1, $2, $3)
RETURNING id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at
`

type UpdateTodoParams struct {
	ActivityGroupID sql.NullInt32
	Title           sql.NullString
	IsActive        sql.NullBool
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) error {
	_, err := q.db.ExecContext(ctx, updateTodo, arg.ActivityGroupID, arg.Title, arg.IsActive)
	return err
}
