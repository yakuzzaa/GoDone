package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ToDoListID  uint `gorm:"index"`
	Title       string
	Description string
	Done        bool
}
