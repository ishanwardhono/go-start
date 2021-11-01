package database

import (
	"app/provider/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func GetDB() *sql.DB {
	if db == nil {
		err := NewDB()
		if err != nil {
			log.Fatal("database initialization failed, err:", err)
		}
		err = db.Ping()
		if err != nil {
			log.Fatal("database connection failed, err:", err)
		}
	}
	return db
}

func NewDB() error {
	cfg := config.GetConfig()
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	newDB, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Println("database error: ", err)
		return err
	}
	db = newDB
	return nil
}
