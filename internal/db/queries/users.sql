-- name: CreateUser :one
INSERT INTO users (email, name)
VALUES ($1, $2)
RETURNING id, email, name, created_at;

-- name: GetUser :one
SELECT id, email, name, created_at FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, name, created_at FROM users ORDER BY id DESC;

-- name: UpdateUser :one
UPDATE users SET email = $2, name = $3 WHERE id = $1
RETURNING id, email, name, created_at;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
