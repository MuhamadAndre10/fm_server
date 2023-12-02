package router

import (
	handler "github.com/andrepriyanto10/server_favaa/internal/user_management/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type ConfigRouter struct {
	App         *fiber.App
	UserHandler *handler.UserHandler
}

func InitRouter(cfg ConfigRouter) {

}
