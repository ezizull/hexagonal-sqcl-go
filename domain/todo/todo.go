package todo

import "database/sql"

type Todo struct {
	ID              int64
	ActivityGroupID sql.NullInt32
	Title           sql.NullString
	IsActive        sql.NullBool
	Priority        sql.NullString
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	DeletedAt       sql.NullTime
}
