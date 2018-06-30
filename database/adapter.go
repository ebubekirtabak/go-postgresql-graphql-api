package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DATABASE_HOST      = "localhost"
	DATABASE_PORT      = 5432
	DATABASE_NAME      = "people"
	DATABASE_USER_NAME = "ebubekirtabak"
	DATABASE_PASSWORD  = ""
	CITIZEN_TABLE_NAME = "citizen"
)

func OpenDatabase() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DATABASE_HOST, DATABASE_PORT, DATABASE_USER_NAME, DATABASE_PASSWORD, DATABASE_NAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestDatabase() {
	db := OpenDatabase()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Successfully connected!")
}
