package users

const (
	userInsertQuery = `
		INSERT INTO 
			users (name, email)
		VALUES
			( $1, $2 )
	`

	userGetAllQuery = `
		SELECT * FROM users
	`

	userGetQuery = `
		SELECT * FROM users
			WHERE name = $1
	`
)
