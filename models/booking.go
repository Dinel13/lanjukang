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
