package handler

import (
	"github.com/andrepriyanto10/server_favaa/internal/user_management"
	"github.com/andrepriyanto10/server_favaa/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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

	// validate user request
	validate := utils.Validate(registerRequest)
	if validate != nil {
		return utils.New(c).Error(validate, fiber.StatusBadRequest)
	}

	err := c.BodyParser(&registerRequest)
	if err != nil {
		return utils.New(c).BadRequest(err.Error())
	}

	var customErr utils.Err

	// call service
	response, err := h.userService.Register(&registerRequest)
	errors.As(err, &customErr)
	if err != nil {
		return utils.New(c).InternalServerError(customErr.Error())
	}

	return utils.New(c).Success("success", fiber.StatusOK, response)
}
