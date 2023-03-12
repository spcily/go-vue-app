package seeds

import "gorm.io/gorm"

func Up(db *gorm.DB) {
	CreateSuperAdmin(db)
}
