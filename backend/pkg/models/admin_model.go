package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;uniqueIndex:idx_admin_username"`
	Password string `gorm:"type:varchar(200);not null"`
	IsActive *bool  `gorm:"type:tinyint(1);default:1;uniqueIndex:idx_admin_username"`
}

func (admin *Admin) BeforeDelete(tx *gorm.DB) error {
	tx.Model(&Admin{}).Where("id = ?", admin.ID).Update("is_active", nil)
	return nil
}
