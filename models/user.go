package models

import "time"

type User struct {
	ID    int    `gorm:"primaryKey"`
	Email string `gorm:"unique;not null"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
