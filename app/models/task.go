package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID       uint
	TaskName string
	Assignee string
	Status   string
	UserID   uint
}

type TaskResponse struct {
	ID       uint
	TaskName string
	Assignee string
	Status   string
	UserID   uint
}
