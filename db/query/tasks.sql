-- name: CreateTask :one
INSERT INTO tasks (
    title, "description", complete
) VALUES (
    $1, $2, $3
) RETURNING *;