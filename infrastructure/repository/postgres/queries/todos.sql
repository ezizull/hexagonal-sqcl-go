

-- name: GetAllTodos :many
SELECT * FROM todos
ORDER BY id DESC;

-- name: GetTodosByActivity :many
SELECT * FROM todos
WHERE activity_group_id = $1
ORDER BY id DESC;

-- name: GetTodosByID :one
SELECT * FROM todos
WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (activity_group_id, title, is_active)
VALUES ($1, $2, $3)
RETURNING *;