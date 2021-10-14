package models

import "time"

type User struct {
	Id        int       `json:"id"`
	FullName  string    `json:"fullname"`
	NickName  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      int       `json:"role"`
	Verified  bool      `json:"verified"`
	Image     string    `json:"image"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserSignUp struct {
	FullName string `json:"fullname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserByEmail struct {
	Email string `json:"email"`
}

type UserById struct {
	Id int `json:"id"`
}

type UserResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

type Service struct {
	Id          int       `json:"id"`
	OwnerId     int       `json:"owner_id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	TypeId      int       `json:"type_id"`
	Capacity    int       `json:"capacity"`
	LocationId  int       `json:"location_id"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	Comments    []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Comment struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Type struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Location struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Image      string    `json:"image"`
	Coordinate string    `json:"coordinate"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

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
