package repo

const (
	articleInsertQuery = `
		INSERT INTO 
			articles (name, email)
		VALUES
			( $1, $2 )
	`

	articleGetAllQuery = `
		SELECT * FROM articles
	`

	articleGetQuery = `
		SELECT * FROM articles
			WHERE name = $1
	`
)
