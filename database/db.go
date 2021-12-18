package database

import (
	"app/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

func GetDB() *sqlx.DB {
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
	var dbinfo string
	if cfg.DBName == "" {
		dbinfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	} else {
		dbinfo = fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBPort, cfg.DBName)
	}
	newDB, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		log.Fatalln("database error: ", err)
		return err
	}
	db = newDB
	return nil
}
