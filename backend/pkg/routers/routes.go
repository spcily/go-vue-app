package routers

import (
	"go-vue-app/backend/pkg/repositories"
	"go-vue-app/backend/pkg/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB, app *fiber.App) {

	adminRepo := repositories.NewAdminRepository(db)
	adminService := services.NewAdminService(adminRepo)
	v1 := app.Group("/v1")
	RegisterAuthRoutes(v1, adminService)
}
