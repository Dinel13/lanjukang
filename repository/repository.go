package repository

import "github.com/dinel13/lanjukang/models"

type DatabaseRepo interface {
	CreateUser(models.UserSignUp) (*models.UserPostSignUp, error)
	GetUserByEmail(string) (*models.UserPostLogin, error)
	GetUserByID(int) (*models.UserById, error)

	CreateService(models.ServiceRequest) (*models.ServicePostCreate, error)
}
