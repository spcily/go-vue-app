package services

import (
	"errors"
	"go-vue-app/backend/app/models"
	"log"
)

func GetAdminByUsername(username string) (*models.Admin, error) {
	for _, admin := range models.Admins {
		log.Println(admin.Username)
		if admin.Username == username {
			return &admin, nil
		}
	}

	return nil, errors.New("admin not found")

}
