// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO Users (
    user_name,
    full_name,
    email, 
    password
) VALUES (
    $1, $2, $3, $4
)
RETURNING user_name, full_name, email, password, created_at, last_updated
`

type CreateUserParams struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserName,
		arg.FullName,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.UserName,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM Users WHERE user_name = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userName string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userName)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_name, full_name, email, password, created_at, last_updated FROM Users WHERE user_name = $1
`

func (q *Queries) GetUser(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userName)
	var i User
	err := row.Scan(
		&i.UserName,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE Users 
SET full_name = $1,
    password = $2
WHERE user_name = $3
`

type UpdateUserParams struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.FullName, arg.Password, arg.UserName)
	return err
}
