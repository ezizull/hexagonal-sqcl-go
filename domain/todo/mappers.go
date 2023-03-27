package todo

func (n *NewTodo) toDomainMapper() *Todo {
	return &Todo{
		Title:           n.Title,
		ActivityGroupID: n.ActivityGroupID,
		IsActive:        n.IsActive,
	}
}
