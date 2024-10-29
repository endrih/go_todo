package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Initialize() *sql.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", host, user, password, dbName)
	db, err := sql.Open("postgres", connectionString)
	checkError(err)
	return db
}
