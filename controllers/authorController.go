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
	author  models.Author
	authors []models.Author
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	query := db.Model(&models.Author{}).Find(&author)
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

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	query := db.Model(&models.Author{}).Where("ID = ?", id).Find(&author)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(author)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	a := &models.Author{}
	utils.ParseBody(r, a)
	query := db.Create(&a)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	a := &models.Author{}
	utils.ParseBody(r, a)

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	query := db.Model(&models.Author{}).Where("ID = ?", id).Find(&author)
	if query.Error != nil {
		panic(query.Error)
	}
	if a.Firstname != "" {
		author.Firstname = a.Firstname
	}
	if a.Lastname != "" {
		author.Lastname = a.Lastname
	}
	query = db.Save(&author)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(author)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	query := db.Model(&models.Author{}).Where("ID = ?", id).Delete(&author)
	if query.Error != nil {
		panic(query.Error)
	}
	res, err := json.Marshal(author)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
