package utils

import "github.com/gofiber/fiber/v2"

// Response is a struct custom Response
type Response struct {
	c *fiber.Ctx
}

func New(c *fiber.Ctx) *Response {
	return &Response{
		c: c,
	}
}

func (r Response) Success(message string, statusCode int, data any) error {
	return r.c.Status(statusCode).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

func (r Response) Error(message any, statusCode int) error {
	return r.c.Status(statusCode).JSON(fiber.Map{
		"message": message,
	})
}

func (r Response) NotFound(message string) error {
	return r.c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": message,
	})
}

func (r Response) BadRequest(message string) error {
	return r.c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
}

func (r Response) InternalServerError(message string) error {
	return r.c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": message,
	})
}

func (r Response) MethodNotAllowed(message string) error {
	return r.c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
		"message": message,
	})
}
