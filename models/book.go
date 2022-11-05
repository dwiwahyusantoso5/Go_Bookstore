package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"not null;type:varchar(100)" json:"title"`
	Author      Author `gorm:"foreignKey:BookID" json:"author"`
	Publication string `gorm:"not null;type:varchar(100)" json:"publication"`
}
