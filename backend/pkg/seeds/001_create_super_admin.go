package seeds

import (
	"go-vue-app/backend/pkg/models"

	"gorm.io/gorm"
)

func CreateSuperAdmin(db *gorm.DB) {
	var admin models.Admin
	db.First(&models.Admin{}, "username = ?", "admin").Scan(&admin)
	if admin.ID == 0 {
		db.Create(&models.Admin{Username: "admin", Password: ""})
	}
}
