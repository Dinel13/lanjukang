package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dinel13/lanjukang/models"
)

// CReateUser creates a new user
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

// Start for all UPDATE USER
//
//
// UpdateUserPasword updates user password
func (m *postgresDbRepo) UpdateUserPasword(id int, password string) (*models.UserPostLogin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE users SET password = $1 WHERE id = $2 RETURNING id, role, nick_name`

	row := m.DB.QueryRowContext(ctx, stmt, password, id)

	var user models.UserPostLogin
	err := row.Scan(
		&user.Id,
		&user.Role,
		&user.NickName,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// BecomeAdmin changes user role to admin
func (m *postgresDbRepo) UpdateUserRole(id int, userData models.UserBecomeAdminRequest) (*models.UserPostLogin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE users SET role = $1, name_service = $2, rekening = $3, bank = $4
				WHERE id = $5
				RETURNING id, role, nick_name`

	row := m.DB.QueryRowContext(ctx, stmt,
		1,
		userData.Name,
		userData.Rekening,
		userData.Bank,
		id,
	)

	var user models.UserPostLogin
	err := row.Scan(
		&user.Id,
		&user.Role,
		&user.NickName,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserProfile updates user profile and returns updated user
func (m *postgresDbRepo) UpdateUserProfile(id int, user models.UserUpdateRequset) (*models.UserDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE users
				SET full_name = $1, nick_name = $2, phone = $3, address = $4 
				WHERE id = $5
				RETURNING id, full_name, nick_name, email, image, phone, address`

	row := m.DB.QueryRowContext(ctx, stmt,
		user.FullName,
		user.NickName,
		user.Phone,
		user.Address,
		id,
	)

	var userUpdated models.UserDetail
	err := row.Scan(
		&userUpdated.Id,
		&userUpdated.FullName,
		&userUpdated.NickName,
		&userUpdated.Email,
		&userUpdated.Image,
		&userUpdated.Phone,
		&userUpdated.Address,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &userUpdated, nil
}

// UpdateUserImage updates user image
func (m *postgresDbRepo) UpdateUserImage(id int, image string) (*models.UserUpdateImage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE users SET image = $1 WHERE id = $2 RETURNING image`

	row := m.DB.QueryRowContext(ctx, stmt, image, id)

	var user models.UserUpdateImage
	err := row.Scan(
		&user.Image,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Start for all GET USER
//
//
// GetUserForUpdateImage returns user info for update image

func (m *postgresDbRepo) GetUserForUpdateImage(id int) (*models.UserUpdateImage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT image FROM users WHERE id = $1`
	row := m.DB.QueryRowContext(ctx, stmt, id)

	var user models.UserUpdateImage
	err := row.Scan(
		&user.Image,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
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

// GetUserForResetPassword returns a user for reset password
func (m *postgresDbRepo) GetUserForResetPassword(id int) (*models.UserForResetPassword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT id, password FROM users WHERE id = $1`
	row := m.DB.QueryRowContext(ctx, stmt, id)

	var user models.UserForResetPassword
	err := row.Scan(
		&user.Id,
		&user.Password,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByNameService returns true if user exists
func (m *postgresDbRepo) GetUserByNameService(name string) (*int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT id FROM users WHERE name_service = $1`
	row := m.DB.QueryRowContext(ctx, stmt, name)

	var idUser int
	err := row.Scan(&idUser)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &idUser, nil
}
