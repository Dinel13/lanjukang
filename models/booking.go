package models

import "time"

type Booking struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	ServiceId     int       `json:"service_id"`
	OwnerId       int       `json:"owner_id"`
	StartAt       time.Time `json:"start_at"`
	EndAt         time.Time `json:"end_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	TransactionId int       `json:"transaction_id"`
	Service       Service   `json:"service"` // supaya bisa ditampilkan di booking
}

// BookingRequest untuk mengambil data booking dari request dan di kirm ke db repo
type BookingRequest struct {
	UserId    int       `json:"user_id"`
	ServiceId int       `json:"service_id"`
	OwnerId   int       `json:"owner_id"`
	Amount    int       `json:"amount"`
	StartAt   time.Time `json:"start_at"`
	EndAt     time.Time `json:"end_at"`
}

// BookingRequestFrontend untuk mengambil data booking dari request fronted dan di kirm handler
type BookingRequestFrontend struct {
	UserId    int    `json:"user_id"`
	ServiceId int    `json:"service_id"`
	OwnerId   int    `json:"owner_id"`
	Amount    int    `json:"amount"`
	StartAt   string `json:"start_at"`
	EndAt     string `json:"end_at"`
}

// BookingRequestUpdate untuk dikirm ke repo DB
type BookingRequestUpdate struct {
	Id      int       `json:"id"`
	UserId  int       `json:"user_id"`
	Amount  int       `json:"amount"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

// BookingRequestUpdateFrontend untuk mengambil data booking dari request fronted dan di kirm handler
type BookingRequestUpdateFrontend struct {
	Id      int    `json:"id"`
	Amount  int    `json:"amount"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}

// BookingRequestDeleteFrontend for delete
type BookingRequestDeleteFrontend struct {
	Id int `json:"id"`
}

// BookingResponse untuk mengiris data booking ke response
type BookingResponse struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	ServiceId int       `json:"service_id"`
	OwnerId   int       `json:"owner_id"`
	Amount    int       `json:"amount"`
	StartAt   time.Time `json:"start_at"`
	EndAt     time.Time `json:"end_at"`
}

type Transaction struct {
	Id     int    `json:"id"`
	Amount int    `json:"amount"`
	Status int    `json:"status"`
	Via    string `json:"via"`
}

type BookingRestriction struct {
	Id         int       `json:"id"`
	BookingId  int       `json:"booking_id"`
	StartAt    time.Time `json:"start_at"`
	EndAt      time.Time `json:"end_at"`
	ServicesId int       `json:"services_id"`
	Name       string    `json:"name"`
}
