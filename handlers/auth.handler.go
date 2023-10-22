package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kurler3/go-task-api/database"
	"github.com/Kurler3/go-task-api/models"
	"github.com/Kurler3/go-task-api/services"
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
	if services.DoesUserExistByEmail(registerBody.Email) {
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
	user := &models.User{
		Name:              registerBody.Name,
		Email:             registerBody.Email,
		EncryptedPassword: hashedPwd,
	}

	// Save user in db
	createUserResult := database.DB.Save(user)

	if createUserResult.Error != nil {
		fmt.Println(createUserResult.Error)
		http.Error(w, "Error while creating user", http.StatusInternalServerError)
		return
	}

	// Create JWT with new user
	jwtToken, err := utils.GenerateToken(*user)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error while generating token", http.StatusInternalServerError)
		return
	}

	// Return JWT to client
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(jwtToken))
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login
func HandleLogin(w http.ResponseWriter, r *http.Request) {

	// Email and password from body
	var loginBody LoginBody

	// Decode body and put in user var
	if err := json.NewDecoder(r.Body).Decode(&loginBody); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get by email
	user, getUserErr := services.GetUserByEmail(loginBody.Email)

	if getUserErr != nil {
		fmt.Println(getUserErr.Error())
		http.Error(w, "Error while getting user", http.StatusInternalServerError)
		return
	}

	// Compare password with encrypted password
	comparePwdErr := utils.ComparePasswords(user.EncryptedPassword, loginBody.Password)

	if comparePwdErr != nil {
		fmt.Println(comparePwdErr.Error())
		http.Error(w, "Permission denied", http.StatusUnauthorized)
		return
	}

	fmt.Println((*user).ID)

	// Create JWT with new user
	jwtToken, err := utils.GenerateToken(*user)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error while generating token", http.StatusInternalServerError)
		return
	}

	// Return JWT to client
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(jwtToken))
}
