package routers

import (
	"go-vue-app/backend/pkg/controllers"
	"go-vue-app/backend/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterAuthRoutes(app fiber.Router, adminService *services.AdminService) {
	authController := controllers.NewAuthController(adminService)

	v1Auth := app.Group("/auth")
	v1Auth.Post("/login", authController.Login)
}
