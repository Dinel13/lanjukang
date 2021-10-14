package models

import "time"

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

// ServiceRequest for incoming request
type ServiceRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	TypeId      int    `json:"type_id"`
	Capacity    int    `json:"capacity"`
	LocationId  int    `json:"location_id"`
	Description string `json:"description"`
}

type ServicePostCreate struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	TypeId      int    `json:"type_id"`
	Capacity    int    `json:"capacity"`
	LocationId  int    `json:"location_id"`
	Description string `json:"description"`
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
