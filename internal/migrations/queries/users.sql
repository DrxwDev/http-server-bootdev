-- name: CreateUser :one
INSERT INTO users (id, email) VALUES ( $1, $2 )
RETURNING *;

-- name: ResetUsers :exec
TRUNCATE TABLE users RESTART IDENTITY CASCADE;
