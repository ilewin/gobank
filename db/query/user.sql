-- name: CreateUser :one 
INSERT INTO users (username, hashed_passwor, full_name, email) VALUES ($1,$2,$3,$4) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users WHERE username = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: GetUsers :many
SELECT * FROM users WHERE full_name = $1 ORDER BY created_at DESC;

-- name: ListUsers :many
SELECT * FROM users ORDER BY username DESC LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users SET full_name=$2, email=$3 WHERE username = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;