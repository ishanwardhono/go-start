package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

const (
	DB_NAME = "ishandb"
)

func NewDB() *sql.DB {
	if db == nil {
		dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
		newDB, err := sql.Open("postgres", dbinfo)
		if err != nil {
			log.Fatal("database error: ", err)
		}
		db = newDB
	}
	return db
}
