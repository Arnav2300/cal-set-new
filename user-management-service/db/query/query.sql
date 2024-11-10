-- name: CreateUserViaEmail :one
INSERT INTO users (email, username, password, role)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUserById :exec
UPDATE users
    set username = $2, 
    password = $3
WHERE id= $1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id= $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at;

-- name: DeleteUserById :exec
DELETE FROM users
WHERE id = $1;