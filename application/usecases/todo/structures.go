package todo

import "database/sql"

type NewTodo struct {
	Title           sql.NullString `json:"title" example:"title todo"`
	ActivityGroupID sql.NullInt32  `json:"activity_group_id" example:"1"`
	IsActive        sql.NullBool   `json:"is_active" example:"false"`
}
