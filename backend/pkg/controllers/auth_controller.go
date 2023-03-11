package controllers

import (
	"errors"
	"go-vue-app/backend/internal/utils"
	"go-vue-app/backend/pkg/response"
	"go-vue-app/backend/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginController(adminService *services.AdminService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var loginDto LoginDto
		if err := c.BodyParser(&loginDto); err != nil {
			return response.ErrorResponse(c, 400, -1, "wrong input", nil)
		}

		admin, err := adminService.Login(loginDto.Username, loginDto.Password)
		if err != nil {
			if errors.Is(err, utils.ErrPasswordValidation) {
				return response.ErrorResponse(c, 400, -1, "wrong password", nil)
			}
			return response.ErrorResponse(c, 400, -1, "login failed", nil)
		}

		return response.SuccessResponse(c, "login success", admin)
	}
}
