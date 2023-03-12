package seeds

import (
	"go-vue-app/backend/pkg/models"

	"gorm.io/gorm"
)

func CreateSuperAdmin(db *gorm.DB) {
	db.FirstOrCreate(&models.Admin{Username: "admin", Password: ""})
}
