package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kurler3/go-task-api/models"
	"github.com/Kurler3/go-task-api/services"
	"github.com/Kurler3/go-task-api/utils"
)

// Create task
func CreateTask(w http.ResponseWriter, r *http.Request) {

	// Declare task var
	var task *models.Task

	// Decode body and check against task struct. "Fill" in task var if ok
	if err := json.NewDecoder(r.Body).Decode(task); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get userID from the authenticated user
	userID := services.GetUserIdFromContext(r)

	task.UserID = userID

	// Create task
	task, err := services.CreateTask(task)

	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	utils.ReturnJSONToClient(
		w,
		*task,
	)
}

// Get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {

	// Get userID from the authenticated user
	userID := services.GetUserIdFromContext(r)

	tasks, err := services.GetUserTasks(userID)

	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*tasks)
}

// Get task by id
func GetTaskById(w http.ResponseWriter, r *http.Request) {

	taskID, err := utils.VarToUint(r, "id")

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Get task
	task, err := services.GetTaskById(taskID)

	// If error
	if err != nil {
		http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		return
	}

	// Get userID from the authenticated user
	userID := services.GetUserIdFromContext(r)

	// Check if task.userID is same as userId
	if task.UserID != userID {
		http.Error(w, "Permission denied", http.StatusUnauthorized)
		return
	}

	// Return to client
	utils.ReturnJSONToClient(w, *task)

}

// Update task
func UpdateTask(w http.ResponseWriter, r *http.Request) {

	taskID, err := utils.VarToUint(r, "id")

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Get task
	task, err := services.GetTaskById(taskID)

	if err != nil {
		http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		return
	}

	// Get userID from the authenticated user
	userID := services.GetUserIdFromContext(r)

	// Check if task.userID is same as userId
	if task.UserID != userID {
		http.Error(w, "Permission denied", http.StatusUnauthorized)
		return
	}

	// Decode the body
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Save task
	task, err = services.UpdateTask(task)

	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	// Return to client
	utils.ReturnJSONToClient(w, *task)
}

// Delete task
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	taskID, err := utils.VarToUint(r, "id")

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Get task
	task, err := services.GetTaskById(taskID)

	// If error
	if err != nil {
		http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		return
	}

	// Get userID from the authenticated user
	userID := services.GetUserIdFromContext(r)

	// If task userId not the same as the userId from the jwt => denied
	if task.UserID != userID {
		http.Error(w, "Permission denied", http.StatusUnauthorized)
		return
	}

	// Delete task
	err = services.DeleteTask(taskID)

	if err != nil {
		http.Error(w, "Error while deleting task", http.StatusBadRequest)
		return
	}

	// Return success msg
	// Send a plain text response message
	message := fmt.Sprintf("Task with ID %d has been deleted", taskID)

	utils.ReturnMessageToClient(w, message, http.StatusOK)

}
