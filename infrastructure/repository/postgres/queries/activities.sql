-- name: GetActivities :many
SELECT * FROM activities
ORDER BY id DESC;