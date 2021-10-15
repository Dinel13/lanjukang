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

// ServiceResponse is a response for list service
// not give detail of service
type ServiceResponse struct {
	Id       int     `json:"id"`
	Owner    string  `json:"owner"`
	OwnerId  int     `json:"owner_id"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	Image    string  `json:"image"`
	Type     string  `json:"type"`
	Capacity int     `json:"capacity"`
	Location string  `json:"location"`
	Rating   float32 `json:"rating"`
}

// ServiceDetailResponse is a responsen for detail service
// return all field of service
type ServiceDetailResponse struct {
	Id          int     `json:"id"`
	Owner       string  `json:"owner"`
	OwnerId     int     `json:"owner_id"`
	Name        string  `json:"name"`
	Price       int     `json:"price"`
	Image       string  `json:"image"`
	Type        string  `json:"type"`
	Capacity    int     `json:"capacity"`
	Location    string  `json:"location"`
	Rating      float32 `json:"rating"`
	Description string  `json:"description"`
	// Comments    []Comment `json:"comments"`
}

// ServiceRequest for incoming request
type ServiceRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	OwnerId     int    `json:"owner_id"`
	TypeId      int    `json:"type_id"`
	Capacity    int    `json:"capacity"`
	LocationId  int    `json:"location_id"`
	Description string `json:"description"`
}

// ServiceUpdateRequest for incoming request
type ServiceUpdateRequest struct {
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

// ServiceDeleteRequest for incoming request delete
type ServiceDeleteRequest struct {
	OwnerId int    `json:"owner_id"`
	Image   string `json:"image"`
}

// ServiceSortDetailResponse  for retunn sort info service
type ServiceSortDetailResponse struct {
	OwnerId int    `json:"owner_id"`
	Image   string `json:"image"`
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
