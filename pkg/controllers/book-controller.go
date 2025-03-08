package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/gorilla/mux"
	"github.com/saatvik-10/bookstore/pkg/models"
	_ "github.com/saatvik-10/bookstore/pkg/utils"
	"net/http"
	_ "strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() //fetch books
	res, _ := json.Marshal(newBooks) //like "application/json" in nodejs
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //200
	_, WriteErr := w.Write(res)  //to send to frontend/postman

	if WriteErr != nil {
		fmt.Println("Error", WriteErr)
	}
}
