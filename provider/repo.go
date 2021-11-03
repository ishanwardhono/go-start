package provider

import (
	"app/provider/database"
	"app/provider/database/repo"
)

func GetUsersRepo() repo.UserRepo {
	return repo.NewUserRepo(
		database.GetDB(),
	)
}
