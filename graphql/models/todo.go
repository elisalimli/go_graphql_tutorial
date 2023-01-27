package models

import (
	"time"
)

type Todo struct {
	ID string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	Name       string `json:"name"`
	IsComplete bool   `json:"isComplete"`
	IsDeleted  bool   `json:"isDeleted"`
	// CreatedId  string
	// User       User `gorm:"foreignKey:CreatedId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TodoInput struct {
	Name      string `json:"name" validate:"required,max=20"`
	CreatedBy int    `json:"user"`
}
