package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

//ConnectDB handle the connection with the database
func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("POSTGRESS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", pgURL)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
