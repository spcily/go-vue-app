package repositories

import (
	"go-vue-app/backend/internal/utils"
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
	// mock api until we have a database
	hash, _ := utils.HashPassword("123456")
	admin.Username = username
	admin.Password = hash

	return &admin, nil

	// err := repo.db.Where("username = ?", username).First(&admin).Error

	// return &admin, err
}
