// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUserViaEmail = `-- name: CreateUserViaEmail :one
INSERT INTO users (id, email, username, password, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, username, password, role, updated_at, created_at
`

type CreateUserViaEmailParams struct {
	ID       pgtype.UUID
	Email    pgtype.Text
	Username string
	Password pgtype.Text
	Role     string
}

func (q *Queries) CreateUserViaEmail(ctx context.Context, arg CreateUserViaEmailParams) (User, error) {
	row := q.db.QueryRow(ctx, createUserViaEmail,
		arg.ID,
		arg.Email,
		arg.Username,
		arg.Password,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUserById = `-- name: DeleteUserById :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserById(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUserById, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, username, password, role, updated_at, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, username, password, role, updated_at, created_at FROM users
WHERE id= $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, email, username, password, role, updated_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, username, password, role, updated_at, created_at FROM users
ORDER BY created_at
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.Role,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserById = `-- name: UpdateUserById :exec
UPDATE users
    set username = $2, 
    password = $3
WHERE id= $1
`

type UpdateUserByIdParams struct {
	ID       pgtype.UUID
	Username string
	Password pgtype.Text
}

func (q *Queries) UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) error {
	_, err := q.db.Exec(ctx, updateUserById, arg.ID, arg.Username, arg.Password)
	return err
}
