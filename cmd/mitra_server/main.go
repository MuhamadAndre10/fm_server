package main

import (
	"github.com/andrepriyanto10/server_favaa/configs/database"
	"github.com/andrepriyanto10/server_favaa/configs/env"
	log "github.com/andrepriyanto10/server_favaa/configs/logger"
	router "github.com/andrepriyanto10/server_favaa/internal"
	handler "github.com/andrepriyanto10/server_favaa/internal/user_management/delivery/http"
	"github.com/andrepriyanto10/server_favaa/internal/user_management/repository"
	"github.com/andrepriyanto10/server_favaa/internal/user_management/service"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"time"
)

//var app *fiber.App

type config struct {
	app *fiber.App
	env *viper.Viper
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.NewLogger("log.txt").ErrorLog.Fatalf("Revocer Panic Error: %v", r)
		}
	}()

	app := initFiber()

	customLog := log.NewLogger("log.txt")

	loadEnv := env.LoadEnv("config", ".")

	conn := database.NewConnection(loadEnv, customLog).Open()

	//

	repoUser := repository.NewUserRepository(conn)

	serviceUser := service.NewUserService(repoUser)
	mailService := service.NewMailService(loadEnv)

	userHandler := handler.NewUMHandler(app, customLog, serviceUser, mailService)

	router.InitRouter(router.ConfigRouter{
		App:         app,
		UserHandler: userHandler,
	})

	go func() {
		err := app.Listen(":" + loadEnv.GetString("APP_PORT"))
		if err != nil {
			customLog.ErrorLog.Panicf("Error starting server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	customLog.InfoLog.Println("Shutting down server...")

	err := app.Shutdown()
	if err != nil {
		customLog.ErrorLog.Fatalf("Error shutting down server: %v", err)
	}

	err = customLog.File.Close()
	if err != nil {
		customLog.InfoLog.Printf("Error closing file: %v", err)
	}

}

func initFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New())

	return app

}
