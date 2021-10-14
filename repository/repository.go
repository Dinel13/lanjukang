package repository

import "github.com/dinel13/lanjukang/models"

type DatabaseRepo interface {
	CreateUser(models.UserSignUp) (int, error)
	GetUserByEmail(string) (*models.UserByEmail, error)
	GetUserByID(int) (*models.UserById, error)
}
