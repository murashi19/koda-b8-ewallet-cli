package models

import (
	"time"
)

type User struct {
	ID          int64
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
