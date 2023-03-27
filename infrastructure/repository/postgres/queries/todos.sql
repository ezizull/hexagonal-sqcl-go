

-- name: GetTodosByActivity :many
SELECT * FROM activities a
JOIN todos t ON a.id = t.activity_group_id
WHERE a.id = $1
ORDER BY id DESC;

-- name: GetSingleTodos :many
SELECT * FROM todos
WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (activity_group_id, title, is_active)
VALUES ($1, $2, $3)
RETURNING *;