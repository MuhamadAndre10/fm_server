package handler

import (
	"fmt"
	"github.com/andrepriyanto10/server_favaa/configs/logger"
	"github.com/andrepriyanto10/server_favaa/internal/user_management"
	"github.com/andrepriyanto10/server_favaa/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	app         *fiber.App
	log         *logger.Log
	userService user_management.UserContractService
	mailService user_management.MailService
}

func NewUMHandler(app *fiber.App, log *logger.Log, user user_management.UserContractService, mail user_management.MailService) *UserHandler {
	return &UserHandler{
		app:         app,
		log:         log,
		userService: user,
		mailService: mail,
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

	// validate user request
	validate := utils.Validate(registerRequest)
	if validate != nil {
		return utils.New(c).Error(validate, fiber.StatusBadRequest)
	}

	code := utils.CodeVerification()

	err = h.userService.Register(&registerRequest, &code)
	if err != nil {
		h.log.InfoLog.Println(fmt.Sprintf("Error when register user: %v", err))
		return utils.New(c).InternalServerError(err.Error())
	}

	data := struct {
		Name string
		Code *string
	}{
		Name: registerRequest.FullName,
		Code: &code,
	}

	dataTmplString, err := utils.ParseTemplate("public/template/email_tmpl.html", data)
	if err != nil {
		h.log.ErrorLog.Println(fmt.Sprintf("Error when parse template: %v", err))
		return err
	}

	err = h.mailService.SendMailWithSmtp([]string{registerRequest.Email}, "Email Verification", dataTmplString)
	if err != nil {
		h.log.ErrorLog.Println(fmt.Sprintf("Error when send email: %v", err))
		return err
	}

	h.log.InfoLog.Println("Email sent")

	return utils.New(c).Success("success", fiber.StatusOK, nil)
}

func (h *UserHandler) VerifyUser(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodPost {
		return utils.New(c).MethodNotAllowed("Method not allowed")
	}

	var code user_management.CodeRequest

	err := c.BodyParser(&code)
	if err != nil {
		return utils.New(c).BadRequest(err.Error())
	}

	validate := utils.Validate(code)
	if validate != nil {
		return utils.New(c).Error(validate, fiber.StatusBadRequest)
	}

	err = h.userService.VerifyUserRegister(&code)
	if err != nil {
		return utils.New(c).InternalServerError(err.Error())
	}

	panic("implement me")

}
