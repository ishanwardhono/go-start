package provider

import (
	"app/database"
	"app/database/repo"
)

func GetUsersRepo() repo.UserRepo {
	return repo.NewUserRepo(
		database.GetDB(),
	)
}
