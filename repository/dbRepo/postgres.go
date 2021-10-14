package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/dinel13/lanjukang/models"
)

func (m *postgresDbRepo) CreateUser(res models.UserSignUp) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int

	stmt := `INSERT INTO users (full_name, nick_name, password, email) VALUES ($1, $2, $3, $4) returning id`

	nickName := res.FullName
	if len(res.FullName) > 8 {
		nickName = res.FullName[:8]
	}

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FullName,
		nickName,
		res.Password,
		res.Email,
	).Scan(&newId)

	if err != nil {
		return 0, err
	}

	return newId, nil

}

// GetUserByEmail returns a user by email
func (m *postgresDbRepo) GetUserByEmail(email string) (*models.UserByEmail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT email FROM users WHERE email = $1`
	row := m.DB.QueryRowContext(ctx, query, email)

	var user models.UserByEmail
	err := row.Scan(&user.Email)
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
