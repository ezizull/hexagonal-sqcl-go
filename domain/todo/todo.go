package todo

import "database/sql"


type Todo struct {
	ID              int64 `json:"id"`
	ActivityGroupID sql.NullInt32 `json:"activity_group_id" example:"1"`
	Title           sql.NullString `json:"title" example:"title todo"`
	IsActive        sql.NullBool `json:"is_active" example:"false"`
	Priority        sql.NullString `json:"priority" example:"very-high"`
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	DeletedAt       sql.NullTime
}



// Service is a interface that contains the methods for the book service
type Service interface {
	Get(int) (*Todo, error)
	GetAll() ([]*Todo, error)
	Create(*Todo) error
	GetByMap(map[string]interface{}) map[string]interface{}
	GetByID(int) (*Todo, error)
	Delete(int) error
	Update(int, map[string]interface{}) (*Todo, error)
}
