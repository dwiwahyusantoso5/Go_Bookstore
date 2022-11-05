package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Firstname string `gorm:"not null" json:"firstname"`
	Lastname  string `gorm:"not null" json:"lastname"`
	BookID    uint
}
