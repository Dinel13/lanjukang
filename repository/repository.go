package repository

import "github.com/dinel13/lanjukang/models"

type DatabaseRepo interface {

	// User
	CreateUser(models.UserSignUp) (*models.UserPostSignUp, error)
	GetUserByEmail(string) (*models.UserPostLogin, error)
	GetUserByID(int) (*models.UserById, error)
	GetUserForResetPassword(int) (*models.UserForResetPassword, error)
	UpdateUserPasword(int, string) (*models.UserPostLogin, error)
	UpdateUserRole(int, models.UserBecomeAdminRequest) (*models.UserPostLogin, error)
	GetUserByNameService(string) (*int, error) // for check if name for become admin is exits
	GetUserForOtherUser(int) (*models.UserDetail, error)
	UpdateUserProfile(int, models.UserUpdateRequset) (*models.UserDetail, error) // not include password and image
	UpdateUserImage(int, string) (*models.UserUpdateImage, error)
	GetUserForUpdateImage(int) (*models.UserUpdateImage, error)

	// Service
	CreateService(models.ServiceRequest) (*models.ServicePostCreate, error)
	UpdateService(int, models.ServiceUpdateRequest) (*models.ServicePostCreate, error)
	DeleteService(int) error
	GetDetailServiceByID(id int) (*models.ServiceDetailResponse, error)
	ListAllServices(limit int) ([]models.ServiceResponse, error)
	ListPopularServices(limit int) ([]models.ServiceResponse, error)
	ListAllServicesByType(typeId int, limit int) ([]models.ServiceResponse, error)
	ListAllServicesByLocation(locationId int, limit int) ([]models.ServiceResponse, error)
	GetSortDetailServiceByID(int) (*models.ServiceSortDetailResponse, error)

	// Booking
	CreateBooking(models.BookingRequest) (*models.BookingResponse, error)
	UpdateBooking(models.BookingRequestUpdate) (*models.BookingResponse, error)
	DeleteBooking(int, int) error
	GetAllBookingByUserId(int) ([]models.BookingResponse, error)
}
