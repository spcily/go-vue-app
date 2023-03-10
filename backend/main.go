package main

import (
	"go-vue-app/backend/app/models"
	"go-vue-app/backend/app/services"
	"go-vue-app/backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app := fiber.New()
	password, err := utils.HashPassword("admin")
	if err == nil {
		admin := models.Admin{
			Username: "admin",
			Password: password,
		}
		models.AddAdmin(admin)
	}

	v1 := app.Group("/v1")
	v1Auth := v1.Group("/auth")

	v1Auth.Post("/login", func(c *fiber.Ctx) error {
		// parse request body into LoginDto
		var loginDto LoginDto
		if err := c.BodyParser(&loginDto); err != nil {
			return c.SendString("Wrong Input")
		}
		admin, err := services.GetAdminByUsername(loginDto.Username)
		if err != nil {
			return c.SendString("user not found")
		}

		isValidated := utils.ValidatePassword(admin.Password, loginDto.Password)
		if !isValidated {
			return c.SendString("wrong password")
		}

		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")

}
