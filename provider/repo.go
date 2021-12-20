package provider

import (
	"app/database"
	"app/database/repo/users"
)

func GetUsersRepo() users.UserRepo {
	return users.NewUserRepo(
		database.GetDB(),
	)
}
