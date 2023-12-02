package handler

import (
	"github.com/andrepriyanto10/server_favaa/internal/user_management"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user_management.UserContractService
	app         *fiber.App
}

func NewUMHandler(userService user_management.UserContractService, app *fiber.App) *UserHandler {
	return &UserHandler{
		userService: userService,
		app:         app,
	}
}
