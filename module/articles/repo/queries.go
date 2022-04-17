package repo

const (
	articleAllColums = `
		id, title, content, author, created_by, created_time, modified_by, modified_time
	`

	articleInsertQuery = `
		INSERT INTO 
			articles (title, content, author, created_by, created_time)
		VALUES
			( :title, :content, :author, :created_by, :created_time )
		RETURNING id
	`

	articleGetAllQuery = `
		SELECT ` + articleAllColums + ` FROM articles
	`

	articleGetQuery = `
		SELECT ` + articleAllColums + ` FROM articles
			WHERE id = $1
	`
)
