package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"not null;type:varchar(191)" json:"title"`
	Author      string `gorm:"not null;type:varchar(191)" json:"author"`
	Publication string `gorm:"not null;type:varchar(191)" json:"publication"`
}
