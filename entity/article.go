package entity

import "app/core/database"

type Article struct {
	database.DbEntity
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}
