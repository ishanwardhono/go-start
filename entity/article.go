package entity

type Article struct {
	DbEntity
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}
