package database

import (
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"database/sql"
)


const (
	DATABASE_HOST      = "localhost"
	DATABASE_PORT      = 5432
	DATABASE_NAME      = "tutorials"
	DATABASE_USER_NAME = "ebubekirtabak"
	DATABASE_PASSWORD  = ""
)

func OpenDatabase() *sql.DB {

	psqlInfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
		 DATABASE_USER_NAME, DATABASE_NAME)
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
