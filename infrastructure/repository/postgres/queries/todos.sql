

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
INSERT INTO todos (activity_group_id, title, is_active, priority)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET 
    title = COALESCE($1, title),
    priority = COALESCE($2, priority),
    is_active = COALESCE($3, is_active),
    updated_at = now()
WHERE id = $4
RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1
RETURNING id;