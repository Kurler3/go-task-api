package services

import (
	"net/http"

	"github.com/Kurler3/go-task-api/database"
	"github.com/Kurler3/go-task-api/models"
)

func DoesUserExistByEmail(email string) bool {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	return result.RowsAffected > 0
}

func DoesUserExistByID(id uint) bool {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return false
	}
	return true
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

func UpdateUser(user *models.User) (*models.User, error) {

	updateUserResult := database.DB.Save(user)

	if updateUserResult.Error != nil {
		return nil, updateUserResult.Error
	}

	return user, nil
}

func DeleteUser(userId uint) error {
	deleteUserResult := database.DB.Delete(&models.User{}, userId)
	if deleteUserResult.Error != nil {
		return deleteUserResult.Error
	}
	return nil
}

func GetUserTasks(userID uint) (*[]models.Task, error) {
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.Tasks, nil
}

// Get userId from context
func GetUserIdFromContext(r *http.Request) uint {
	return r.Context().Value("userId").(uint)
}
