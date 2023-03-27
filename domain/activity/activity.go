package activity

import "database/sql"

type Activity struct {
	Int       int64
	Title     sql.NullString
	Email     sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
