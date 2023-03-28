package todo

import (
	"time"
)

type Todo struct {
	ID              int64  `json:"id"`
	ActivityGroupID int32  `json:"activity_group_id" example:"1"`
	Title           string `json:"title" example:"title todo"`
	IsActive        bool   `json:"is_active" example:"false"`
	Priority        string `json:"priority" example:"very-high"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type NewTodo struct {
	ActivityGroupID *int32  `json:"activity_group_id" example:"1"`
	Title           *string `json:"title" example:"title todo"`
	IsActive        *bool   `json:"is_active" example:"false"`
}

type UpdateTodo struct {
	Title    *string `json:"title,omitempty" example:"title todo"`
	Priority *string `json:"priority,omitempty" example:"very-high"`
	IsActive *bool   `json:"is_active,omitempty" example:"false"`
}
