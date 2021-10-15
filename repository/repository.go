package repository

import "github.com/dinel13/lanjukang/models"

type DatabaseRepo interface {
	CreateUser(models.UserSignUp) (*models.UserPostSignUp, error)
	GetUserByEmail(string) (*models.UserPostLogin, error)
	GetUserByID(int) (*models.UserById, error)
	UpdateUserRole(int) error

	CreateService(models.ServiceRequest) (*models.ServicePostCreate, error)
	UpdateService(int, models.ServiceUpdateRequest) (*models.ServicePostCreate, error)

	GetDetailServiceByID(id int) (*models.ServiceDetailResponse, error)

	ListAllServices(limit int) ([]models.ServiceResponse, error)
	ListAllServicesByType(typeId int, limit int) ([]models.ServiceResponse, error)
	ListAllServicesByLocation(locationId int, limit int) ([]models.ServiceResponse, error)
}
