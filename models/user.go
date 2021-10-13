package models

type User struct {
	Id       int64  `json:"id"`
	FullName string `json:"fullname"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
