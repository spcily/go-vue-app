package services

import (
	"go-vue-app/backend/internal/utils"
	"go-vue-app/backend/pkg/models"
	"go-vue-app/backend/pkg/repositories"
)

type AdminService struct {
	adminRepo *repositories.AdminRepository
}

func NewAdminService(adminRepo *repositories.AdminRepository) *AdminService {
	return &AdminService{adminRepo}
}

func (service *AdminService) Login(username string, password string) (*models.Admin, error) {
	admin, err := service.adminRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	err = utils.ValidatePassword(admin.Password, password)
	if err != nil {
		return nil, err
	}

	return admin, err
}
