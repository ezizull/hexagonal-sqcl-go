package activity

import (
	"time"
)

type Activity struct {
	ID        int64  `json:"id"`
	Title     string `json:"title" example:"title activity"`
	Email     string `json:"email" example:"activity@email.com"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewActivity struct {
	Title *string `json:"title" example:"title activity"`
	Email *string `json:"email" example:"activity@email.com"`
}

type UpdateActivity struct {
	Title *string `json:"title" example:"title activity"`
}
