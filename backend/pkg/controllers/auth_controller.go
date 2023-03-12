package controllers

import (
	"errors"
	"go-vue-app/backend/internal/utils"
	"go-vue-app/backend/pkg/response"
	"go-vue-app/backend/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AdminService *services.AdminService
}

func NewAuthController(adminService *services.AdminService) *AuthController {
	return &AuthController{AdminService: adminService}
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (authController *AuthController) Login(c *fiber.Ctx) error {
	var loginDto LoginDto
	if err := c.BodyParser(&loginDto); err != nil {
		return response.ErrorResponse(c, 400, -1, "wrong input", nil)
	}

	admin, err := authController.AdminService.Login(loginDto.Username, loginDto.Password)
	if err != nil {
		if errors.Is(err, utils.ErrPasswordValidation) {
			return response.ErrorResponse(c, 400, -1, "wrong password", nil)
		}
		return response.ErrorResponse(c, 400, -1, "login failed", nil)
	}

	return response.SuccessResponse(c, "login success", admin)
}
