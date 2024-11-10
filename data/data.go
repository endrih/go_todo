package data

import (
	"database/sql"
	"endrih/go_todo/config"
	"fmt"

	_ "github.com/lib/pq"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Initialize(config *config.DbSettings) *sql.DB {
	host := config.DB_HOST
	user := config.DB_USER
	password := config.DB_PASSWORD
	dbName := config.DB_NAME
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", host, user, password, dbName)
	db, err := sql.Open("postgres", connectionString)
	checkError(err)
	return db
}
