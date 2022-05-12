package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID        uint
	TaskName  string
	Assignee  string
	Status    string
	CreatedAt time.Time
}
