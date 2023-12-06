package router

import (
	"fmt"
	handler "github.com/andrepriyanto10/server_favaa/internal/user_management/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type ConfigRouter struct {
	App         *fiber.App
	UserHandler *handler.UserHandler
}

func InitRouter(cfg ConfigRouter) {

	micro := fiber.New()

	cfg.App.Mount("/api", micro)

	micro.Get("/test-health", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", cfg.UserHandler.Register)
		router.Post("/verify", cfg.UserHandler.VerifyUser)
		router.Post("/login", cfg.UserHandler.Login)
		router.Post("/get-otp", cfg.UserHandler.ForgotPassword)
	})

	micro.Use(func(c *fiber.Ctx) error {
		path := c.Path() // => "/api/hello"
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		}) // => 404 "Not Found"
	})
}
