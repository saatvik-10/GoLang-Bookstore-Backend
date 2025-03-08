package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/saatvik-10/bookstore/pkg/models"
	"github.com/saatvik-10/bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter) {
	newBooks := models.GetAllBooks() //fetch books
	res, _ := json.Marshal(newBooks) //like "application/json" in nodejs
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //200
	_, WriteErr := w.Write(res)  //to send to frontend/postman

	if WriteErr != nil {
		fmt.Println("Error", WriteErr)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing", err)
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, WriteErr := w.Write(res) //to send to frontend/postman

	if WriteErr != nil {
		fmt.Println("Error", WriteErr)
	}
}
