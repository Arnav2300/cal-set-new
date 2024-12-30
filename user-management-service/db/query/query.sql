-- name: CreateUserViaEmail :one
INSERT INTO users (id, email, username, password, role)
VALUES ($1, $2, $3, $4, $5)
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

-- name: CreatePasswordResetToken :exec
INSERT INTO password_reset_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3);

SELECT COUNT(*) > 0 AS is_valid
FROM password_reset_tokens
WHERE token = $1 AND expires_at > NOW();

-- name: GetPasswordResetTokenByUserID :one
SELECT user_id, token, expires_at, created_at
FROM password_reset_tokens
WHERE user_id = $1;

-- name: GetPasswordResetToken :one
SELECT user_id, token, expires_at, created_at
FROM password_reset_tokens
WHERE token = $1;

-- name: DeletePasswordResetToken :exec
DELETE FROM password_reset_tokens
WHERE user_id = $1;

-- name: DeleteExpiredTokens :exec
DELETE FROM password_reset_tokens
WHERE expires_at <= NOW();