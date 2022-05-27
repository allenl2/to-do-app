package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Tasks    []Task `json:"tasks"`
}
