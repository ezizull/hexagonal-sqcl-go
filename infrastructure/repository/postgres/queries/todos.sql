

-- name: GetTodosByActivity :many
SELECT *
FROM todos
WHERE ($1::int IS NULL OR activity_group_id = $1)
ORDER BY id DESC;

-- name: GetSingleTodos :many
SELECT * FROM todos
WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (activity_group_id, title, is_active)
VALUES ($1, $2, $3)
RETURNING *;