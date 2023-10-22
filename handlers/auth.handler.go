package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kurler3/go-task-api/database"
	"github.com/Kurler3/go-task-api/models"
	"github.com/Kurler3/go-task-api/utils"
)

// Register body
type RegisterBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register
func HandleRegister(w http.ResponseWriter, r *http.Request) {

	// Declare user variable
	var registerBody RegisterBody

	// Decode body and put in user var
	if err := json.NewDecoder(r.Body).Decode(&registerBody); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check that user doesn't already exist
	if utils.DoesUserExistByEmail(registerBody.Email) {
		http.Error(w, "Permission denied", http.StatusForbidden)
		return
	}

	// Get hashed pwd
	hashedPwd, err := utils.EncryptPassword(registerBody.Password)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Encryption error", http.StatusInternalServerError)
		return
	}

	// Create user
	user := models.User{
		Name:              registerBody.Name,
		Email:             registerBody.Email,
		EncryptedPassword: hashedPwd,
	}

	// Save user in db
	createUserResult := database.DB.Save(&user)

	if createUserResult.Error != nil {
		fmt.Println(createUserResult.Error)
		http.Error(w, "Error while creating user", http.StatusInternalServerError)
		return
	}

	// Create JWT with new user

	// Return JWT to client

}

// Login
func HandleLogin(w http.ResponseWriter, r *http.Request) {

	// Email and password from body

	// Check that user exists

	// Compare password with encrypted password

	// Create JWT

	// Return JWT

}
