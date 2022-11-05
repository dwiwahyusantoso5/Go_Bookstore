package controllers

import (
	"encoding/json"
	"go-bookstore/database"
	"go-bookstore/models"
	"go-bookstore/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	book  models.Book
	books []models.Book
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	query := db.Model(&models.Book{}).Preload("Author").Find(&books)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	query := db.Model(&models.Book{}).Preload("Author").Where("ID = ?", id).Find(&book)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	b := &models.Book{}
	utils.ParseBody(r, b)
	query := db.Create(&b)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	b := &models.Book{}
	utils.ParseBody(r, b)

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	query := db.Model(&models.Book{}).Where("ID = ?", id).Find(&book)
	if query.Error != nil {
		panic(query.Error)
	}
	if b.Title != "" {
		book.Title = b.Title
	}
	if b.Author != author {
		book.Author = b.Author
	}
	if b.Publication != "" {
		book.Publication = b.Publication
	}
	query = db.Save(&book)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	query := db.Model(&models.Book{}).Where("ID = ?", id).Delete(&book)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
