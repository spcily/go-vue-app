package repositories

import (
	"go-vue-app/backend/pkg/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db}
}

func (repo *AdminRepository) FindByUsername(username string) (*models.Admin, error) {
	var admin models.Admin

	err := repo.db.Where("username = ?", username).First(&admin).Error

	return &admin, err
}
