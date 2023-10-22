package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kurler3/go-task-api/models"
	"github.com/Kurler3/go-task-api/services"
	"github.com/Kurler3/go-task-api/utils"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {

	// Get user id from token
	userID := services.GetUserIdFromContext(r)

	// Get user from db
	user, err := services.GetUserById(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return to client
	utils.ReturnJSONToClient(w, *user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Declare user var
	var user *models.User

	// Decode body and check against task struct. "Fill" in task var if ok
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get userId from token
	userID := services.GetUserIdFromContext(r)

	// Set userId
	user.ID = userID

	// Update user in db
	user, err := services.UpdateUser(
		user,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return to client
	utils.ReturnJSONToClient(w, *user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// Get userId from token
	userID := services.GetUserIdFromContext(r)

	// Check if user exists
	if userExists := services.DoesUserExistByID(userID); !userExists {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	// Delete user from db
	err := services.DeleteUser(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create msg
	message := fmt.Sprintf("User with id %d deleted", userID)

	// Return msg
	utils.ReturnMessageToClient(
		w,
		message,
		http.StatusOK,
	)
}
