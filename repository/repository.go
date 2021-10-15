package repository

import "github.com/dinel13/lanjukang/models"

type DatabaseRepo interface {

	// User
	CreateUser(models.UserSignUp) (*models.UserPostSignUp, error)
	GetUserByEmail(string) (*models.UserPostLogin, error)
	GetUserByID(int) (*models.UserById, error)
	UpdateUserRole(int) error
	GetUserForOtherUser(int) (*models.UserDetail, error)
	UpdateUserProfile(int, models.UserUpdateRequset) (*models.UserDetail, error) // not include password and image

	// Service
	CreateService(models.ServiceRequest) (*models.ServicePostCreate, error)
	UpdateService(int, models.ServiceUpdateRequest) (*models.ServicePostCreate, error)
	DeleteService(int) error
	GetDetailServiceByID(id int) (*models.ServiceDetailResponse, error)
	ListAllServices(limit int) ([]models.ServiceResponse, error)
	ListAllServicesByType(typeId int, limit int) ([]models.ServiceResponse, error)
	ListAllServicesByLocation(locationId int, limit int) ([]models.ServiceResponse, error)
	GetSortDetailServiceByID(int) (*models.ServiceSortDetailResponse, error)
}
