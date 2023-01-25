package models

import "time"

type Todo struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `json:"name"`
	IsComplete bool   `json:"isComplete"`
	IsDeleted  bool   `json:"isDeleted"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TodoInput struct {
	Name      string `json:"name"`
	CreatedBy int    `json:"user"`
}
