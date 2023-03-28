// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: activities.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createActivity = `-- name: CreateActivity :one
INSERT INTO activities (title, email)
VALUES ($1, $2)
RETURNING id, title, email, created_at, updated_at, deleted_at
`

type CreateActivityParams struct {
	Title sql.NullString
	Email sql.NullString
}

func (q *Queries) CreateActivity(ctx context.Context, arg CreateActivityParams) (Activity, error) {
	row := q.db.QueryRowContext(ctx, createActivity, arg.Title, arg.Email)
	var i Activity
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getActivities = `-- name: GetActivities :many
SELECT id, title, email, created_at, updated_at, deleted_at FROM activities
ORDER BY id DESC
`

func (q *Queries) GetActivities(ctx context.Context) ([]Activity, error) {
	rows, err := q.db.QueryContext(ctx, getActivities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Activity
	for rows.Next() {
		var i Activity
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Email,
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

const getActivityByID = `-- name: GetActivityByID :one
SELECT id, title, email, created_at, updated_at, deleted_at FROM activities
WHERE id = $1
ORDER BY id DESC
`

func (q *Queries) GetActivityByID(ctx context.Context, id int64) (Activity, error) {
	row := q.db.QueryRowContext(ctx, getActivityByID, id)
	var i Activity
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
