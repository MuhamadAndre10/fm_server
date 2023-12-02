package main

import (
	"github.com/andrepriyanto10/server_favaa/configs/database"
	"github.com/andrepriyanto10/server_favaa/configs/env"
	localLogger "github.com/andrepriyanto10/server_favaa/configs/logger"
	router "github.com/andrepriyanto10/server_favaa/internal"
	handler "github.com/andrepriyanto10/server_favaa/internal/user_management/delivery/http"
	"github.com/andrepriyanto10/server_favaa/internal/user_management/repository"
	"github.com/andrepriyanto10/server_favaa/internal/user_management/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

func run() {
	serve := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	serve.Use(logger.New(logger.Config{
		Format: "${pid} ${status} - ${method} ${path}\n",
	}))

	serve.Use(cors.New())

	log := localLogger.NewLogger()

	loadEnv := env.LoadEnv("config", "../../")

	conn := database.NewConnection(loadEnv, log).Open()

	repoUser := repository.NewUserRepository(conn)

	serviceUser := service.NewUserService(repoUser)

	userHandler := handler.NewUMHandler(serviceUser, serve)

	router.InitRouter(router.ConfigRouter{
		App:         serve,
		UserHandler: userHandler,
	})

	app = serve

}
