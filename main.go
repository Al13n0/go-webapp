package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

//load env variables
func init() {
	gotenv.Load()
}

func main() {

	//connect to DB
	db = driver.ConnectDB()

	//initialize the controller
	controller := controllers.Controllers{}

	router := mux.NewRouter()
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook(db)).Methods("DELETE")

	fmt.Print("server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
