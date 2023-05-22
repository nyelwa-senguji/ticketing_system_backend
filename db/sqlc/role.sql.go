// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: role.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createRole = `-- name: CreateRole :execresult
INSERT INTO roles (
  role_name, status, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
)
`

type CreateRoleParams struct {
	RoleName  string    `json:"role_name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createRole,
		arg.RoleName,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const getRole = `-- name: GetRole :one
SELECT id, role_name, status, updated_at, created_at FROM roles
WHERE id = ? LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, id int32) (Roles, error) {
	row := q.db.QueryRowContext(ctx, getRole, id)
	var i Roles
	err := row.Scan(
		&i.ID,
		&i.RoleName,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listRoles = `-- name: ListRoles :many
SELECT id, role_name, status, updated_at, created_at FROM roles
ORDER BY role_name
`

func (q *Queries) ListRoles(ctx context.Context) ([]Roles, error) {
	rows, err := q.db.QueryContext(ctx, listRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Roles
	for rows.Next() {
		var i Roles
		if err := rows.Scan(
			&i.ID,
			&i.RoleName,
			&i.Status,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRole = `-- name: UpdateRole :exec
UPDATE roles
SET role_name=?, status=?, updated_at=?
WHERE id=?
`

type UpdateRoleParams struct {
	RoleName  string    `json:"role_name"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        int32     `json:"id"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.ExecContext(ctx, updateRole,
		arg.RoleName,
		arg.Status,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
