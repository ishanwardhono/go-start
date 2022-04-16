package database

import (
	"app/core/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	dbInstance *DB
)

type DB struct {
	db *sqlx.DB
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
	if cfg.DBUser != "" {
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

//Entity Interface
type IDbEntity interface {
	Insert(user string)
	Update(user string)
}

type DbEntity struct {
	ID           int       `db:"id"`
	CreatedBy    string    `db:"created_by"`
	CreatedTime  time.Time `db:"created_time"`
	ModifiedBy   string    `db:"modified_by"`
	ModifiedTime time.Time `db:"modified_time"`
}

func (e *DbEntity) Insert(user string) {
	e.CreatedTime = time.Now()
	e.CreatedBy = user
}

func (e *DbEntity) Update(user string) {
	e.ModifiedBy = user
	e.ModifiedTime = time.Now()
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
