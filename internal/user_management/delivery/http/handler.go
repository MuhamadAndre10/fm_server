package handler

import (
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/andrepriyanto10/server_favaa/configs/logger"
	"github.com/andrepriyanto10/server_favaa/internal/user_management"
	"github.com/andrepriyanto10/server_favaa/pkg/cache"
	"github.com/andrepriyanto10/server_favaa/utils"
	"github.com/gofiber/fiber/v2"
	"time"
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

	err = h.userService.Register(c.Context(), &registerRequest, &code)
	if err != nil {
		h.log.InfoLog.Println(fmt.Sprintf("Error when register user: %v", err))
		return utils.New(c).InternalServerError(err.Error())
	}

	bigCache, err := bigcache.New(c.Context(), bigcache.DefaultConfig(5*time.Minute))
	if err != nil {
		return err
	}

	newDataCache := struct {
		Email     string
		Code      *string
		ExpiredAt time.Time
	}{
		Email:     registerRequest.Email,
		Code:      &code,
		ExpiredAt: time.Now().Add(60 * time.Second),
	}

	dataByte, err := json.Marshal(newDataCache)
	if err != nil {
		return err
	}

	dataCache := cache.NewDataCache(bigCache)

	err = dataCache.Set("user", dataByte)
	if err != nil {
		return err
	}

	cache.NewCache(dataCache)

	data := struct {
		Name string
		Code *string
	}{
		Name: registerRequest.FirstName,
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

	defer func() {
		if r := recover(); r != nil {
			h.log.ErrorLog.Fatalf(fmt.Sprintf("Panic: %v", r))
		}
	}()

	var code user_management.CodeRequest

	err := c.BodyParser(&code)
	if err != nil {
		return utils.New(c).BadRequest(err.Error())
	}

	validate := utils.Validate(code)
	if validate != nil {
		return utils.New(c).Error(validate, fiber.StatusBadRequest)
	}

	err = h.userService.VerifyUserRegister(c.Context(), &code)
	if err != nil {
		h.log.ErrorLog.Println(fmt.Sprintf("Service Error: %v", err))
		return utils.New(c).InternalServerError(err.Error())
	}

	return utils.New(c).Success("success", fiber.StatusOK, nil)

}
