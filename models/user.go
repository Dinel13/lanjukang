package models

import "time"

type User struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	NickName  string    `json:"nick_name"`
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

// UserSignUp is for reusetsignup
type UserSignUp struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserPOstSignUp is result from query db Signup
type UserPostSignUp struct {
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
}

// UserLogin is for login request
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserPOstLOgin is result from query db login
type UserPostLogin struct {
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	Role     int    `json:"role"`
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
