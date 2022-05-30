package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID       uint
	TaskName string
	Assignee string
	IsDone   bool `gorm:"default:false" json:"isDone"`
	UserID   uint
}

type TaskResponse struct {
	ID       uint
	TaskName string
	Assignee string
	IsDone   bool `gorm:"default:false" json:"isDone"`
	UserID   uint
}
