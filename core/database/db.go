package database

import (
	"app/core/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/guregu/null.v3"
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
	ID           int         `json:"id" db:"id"`
	CreatedBy    string      `json:"-" db:"created_by"`
	CreatedTime  time.Time   `json:"-" db:"created_time"`
	ModifiedBy   null.String `json:"-" db:"modified_by"`
	ModifiedTime null.Time   `json:"-" db:"modified_time"`
}

func (e *DbEntity) Insert(user string) {
	e.CreatedTime = time.Now()
	e.CreatedBy = user
}

func (e *DbEntity) Update(user string) {
	e.ModifiedBy = null.NewString(user, true)
	e.ModifiedTime = null.NewTime(time.Now(), true)
}

func (i *DB) NamedQueryContext(ctx context.Context, query string, entity IDbEntity) (*sqlx.Rows, error) {
	entity.Insert("")
	return i.db.NamedQueryContext(ctx, query, entity)
}

func (i *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return i.db.GetContext(ctx, dest, query, args...)
}

func (i *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return i.db.SelectContext(ctx, dest, query, args...)
}
