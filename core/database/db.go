package database

import (
	"app/core/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	dbInstance *DB
)

type DB struct {
	db *sqlx.DB
}

type IDbEntity interface {
	Insert(user string)
	Update(user string)
}

func GetDB() *DB {
	if dbInstance == nil {
		err := NewDB()
		if err != nil {
			log.Fatal("database initialization failed, err:", err)
		}
		err = dbInstance.db.Ping()
		if err != nil {
			log.Fatal("database connection failed, err:", err)
		}
	}
	return dbInstance
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
	dbInstance = &DB{
		db: newDB,
	}
	return nil
}

func (i *DB) NamedExec(query string, entity IDbEntity) (sql.Result, error) {
	entity.Insert("")
	return i.db.NamedExec(query, entity)
}

func (i *DB) Get(dest interface{}, query string, args ...interface{}) error {
	return i.db.Get(dest, query, args...)
}

func (i *DB) Select(dest interface{}, query string, args ...interface{}) error {
	return i.db.Select(dest, query, args...)
}
