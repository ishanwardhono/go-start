package provider

import (
	"app/handler"
)

func GetHandlers() []handler.HttpHandler {
	return []handler.HttpHandler{
		GetUserHandler(),
	}
}

func GetUserHandler() handler.HttpHandler {
	return handler.NewUserHandler(
		GetUserFactory(),
	)
}
