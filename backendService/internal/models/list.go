package models

import "gorm.io/gorm"

type ToDoList struct {
	gorm.Model
	UserID      uint `gorm:"index"`
	Title       string
	Description string
	Items       []Item `gorm:"foreignKey:ToDoListID"`
}
