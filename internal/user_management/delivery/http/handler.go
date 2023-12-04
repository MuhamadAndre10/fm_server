package handler

import (
	"github.com/andrepriyanto10/server_favaa/internal/user_management"
	"github.com/andrepriyanto10/server_favaa/utils"
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

// Register Handler to register new user
func (h *UserHandler) Register(c *fiber.Ctx) error {
	// check if request method is POST
	if c.Method() != fiber.MethodPost {
		return utils.New(c).MethodNotAllowed("Method not allowed")
	}

	var registerRequest user_management.UserRegisterRequest

	err := c.BodyParser(&registerRequest)
	if err != nil {
		return utils.New(c).BadRequest(err.Error())
	}

	// call service
	_, err = h.userService.Register(&registerRequest)
	if err != nil {
		return err
	}

	panic("implement me")
}
