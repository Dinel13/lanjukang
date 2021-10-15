package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/dinel13/lanjukang/models"
)

func (m *postgresDbRepo) CreateUser(res models.UserSignUp) (*models.UserPostSignUp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.UserPostSignUp

	stmt := `INSERT INTO users (full_name, nick_name, password, email) VALUES ($1, $2, $3, $4) returning id, nick_name, role`

	nickName := res.FullName
	if len(res.FullName) > 8 {
		nickName = res.FullName[:8]
	}

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FullName,
		nickName,
		res.Password,
		res.Email,
	).Scan(
		&user.Id,
		&user.NickName,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// GetUserByEmail returns a user by email
func (m *postgresDbRepo) GetUserByEmail(email string) (*models.UserPostLogin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, nick_name, password, role FROM users WHERE email = $1`
	row := m.DB.QueryRowContext(ctx, query, email)

	var user models.UserPostLogin
	err := row.Scan(
		&user.Id,
		&user.NickName,
		&user.Password,
		&user.Role,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID returns a user by id
func (m *postgresDbRepo) GetUserByID(id int) (*models.UserById, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT id FROM users WHERE id = $1`
	row := m.DB.QueryRowContext(ctx, stmt, id)

	var user models.UserById
	err := row.Scan(&user.Id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// BecomeAdmin changes user role to admin
func (m *postgresDbRepo) UpdateUserRole(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE users SET role = $1 WHERE id = $2`
	_, err := m.DB.ExecContext(ctx, stmt, 1, id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserForOtherUser returns a user info for other user
func (m *postgresDbRepo) GetUserForOtherUser(id int) (*models.UserDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT id, full_name, nick_name, email, image, phone, address FROM users WHERE id = $1`
	row := m.DB.QueryRowContext(ctx, stmt, id)

	var user models.UserDetail
	err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.NickName,
		&user.Email,
		&user.Image,
		&user.Phone,
		&user.Address,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
