package entity

import "app/core/database"

type Article struct {
	database.DbEntity
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Author  string `json:"author" db:"author"`
}
