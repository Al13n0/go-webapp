package controllers

import (
	"books-list/models"
	"books-list/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Controllers struct{}

var books []models.Book

//GetBooks all the books inside the books tables
func (c Controllers) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var books = []models.Book{}
		rows, err := db.Query("select * from books")

		if err != nil {
			log.Fatal("error")
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			if err != nil {
				log.Fatal(err)
			}
			books = append(books, book)
		}
		json.NewEncoder(w).Encode(books)

	}
}

//GetBook get a single book record
func (c Controllers) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		//read parameters from book
		params := mux.Vars(r)

		rows := db.QueryRow("select * from books where id=$1", params["id"])
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Book Not found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}

		}
		json.NewEncoder(w).Encode(book)
		log.Print("getBook was called")
	}
}

//AddBook into books tables
func (c Controllers) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)

		err := db.QueryRow("insert into books (title,author,year) values ($1,$2,$3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, bookID)
		//json.NewEncoder(w).Encode(bookID)
	}
}

//UpdateBook function to update book record
func (c Controllers) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All Fields are required"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)

		rowUpdated, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(rowUpdated)

	}
}

//DeleteBook delete a book record based on id
func (c Controllers) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//params contains the id of the book to remove
		params := mux.Vars(r)

		var book models.Book
		var error models.Error
		json.NewDecoder(r.Body).Decode(&book)

		result, err := db.Exec("delete from books where id = $1", params["id"])

		if err != nil {
			error.Message = "server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return

		}

		rowDeleted, err := result.RowsAffected()
		if err != nil {
			error.Message = "server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		if rowDeleted == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		json.NewEncoder(w).Encode(rowDeleted)

	}
}
