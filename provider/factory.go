package provider

import "app/module/users"

func GetUserFactory() users.Factory {
	return users.NewUsersFactory(
		GetUsersRepo(),
	)
}
