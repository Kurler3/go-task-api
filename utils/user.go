package utils

import (
	"github.com/Kurler3/go-task-api/database"
	"github.com/Kurler3/go-task-api/models"
)

func DoesUserExistByEmail(email string) bool {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	return result.RowsAffected > 0
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	getUserByEmailResult := database.DB.Where("email = ?", email).First(&user)
	if getUserByEmailResult.Error != nil {
		return nil, getUserByEmailResult.Error
	}
	return &user, nil
}

func GetUserById(id uint) (*models.User, error) {
	var user models.User
	getUserByIdResult := database.DB.First(&user, id)
	if getUserByIdResult.Error != nil {
		return nil, getUserByIdResult.Error
	}
	return &user, nil
}
