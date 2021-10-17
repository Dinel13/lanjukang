package models

import (
	"time"
)

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

// UserRespon for respon to frontend
type UserResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

// UserDetail is for user detail
// this will use to other user look thi user
type UserDetail struct {
	Id       int     `json:"id"`
	FullName string  `json:"full_name"`
	NickName string  `json:"nick_name"`
	Email    string  `json:"email"`
	Image    *string `json:"image"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}

// UserUpdateRequset is for request update user
type UserUpdateRequset struct {
	FullName string  `json:"full_name"`
	NickName string  `json:"nick_name"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}

// UserResetPasswordRequerst is for request reset password
type UserResetPasswordRequest struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

// UserForResetPassword is for user for reset password
type UserForResetPassword struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

// UserUpdateImage is for update user image
type UserUpdateImage struct {
	Image *string `json:"image"`
}

// // UserForGetUpdateImage is for user for reset password
// type UserForUpdateImage struct {
// 	Id    int     `json:"id"`
// 	Image *string `json:"image"`
// }
