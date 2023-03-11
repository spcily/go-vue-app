package response

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(200).JSON(fiber.Map{
		"code":    0,
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, errorCode int, message string, err interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"code":    errorCode,
		"message": message,
		"error":   err,
	})
}
