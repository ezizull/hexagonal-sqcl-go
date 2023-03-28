-- name: GetActivities :many
SELECT * FROM activities
ORDER BY id DESC;

-- name: GetActivityByID :one
SELECT * FROM activities
WHERE id = $1
ORDER BY id DESC;

-- name: CreateActivity :one
INSERT INTO activities (title, email)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateActivity :one
UPDATE activities
SET 
    title = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;