package models

type LoginRequest struct {
	Email    string
	Password string
}

var CurrentUser *User

type RegisterRequest struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
}
