package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"required,min=5,max=16,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=20,alphanum"`
	Tasks    []Task `json:"tasks"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Tasks    []Task `json:"tasks"`
}
