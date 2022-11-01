package database

import (
	"fmt"
	"go-bookstore/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	configInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(configInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Debug().AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
