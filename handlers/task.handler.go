package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kurler3/go-task-api/database"
	"github.com/Kurler3/go-task-api/models"
	"github.com/Kurler3/go-task-api/utils"
)

// Create task
func CreateTask(w http.ResponseWriter, r *http.Request) {

	// Declare task var
	var task models.Task

	// Decode body and check against task struct. "Fill" in task var if ok
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get userID from the authenticated user

	userID := uint(1)

	task.UserID = userID

	result := database.DB.Create(&task)
	if result.Error != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Get userID from the authenticated user (you will implement this part later with authentication)
	userID := uint(1) // Example userID

	var user models.User

	findUserResult := database.DB.First(&user, userID)

	if findUserResult.Error != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Tasks)
}

// Get task by id
func GetTaskById(w http.ResponseWriter, r *http.Request) {

	taskID, err := utils.VarToUint(r, "id")

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Declare task
	var task models.Task

	// Get task by id
	getTaskResult := database.DB.First(&task, taskID)

	// If error
	if getTaskResult.Error != nil {
		http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		return
	}

	//TODO Check that the task.userId is same as the id of the user from JWT

	// Return the task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Update task
func UpdateTask(w http.ResponseWriter, r *http.Request) {

	taskID, err := utils.VarToUint(r, "id")

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Declare task
	var task models.Task

	// Get task
	getTaskResult := database.DB.First(&task, taskID)

	if getTaskResult.Error != nil {
		http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		return
	}

	//TODO Check if userId of task is same as id of user from JWT

	// Decode the body
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Save task
	updateTaskResult := database.DB.Save(&task)

	if updateTaskResult.Error != nil {

		http.Error(w, "Error while updating task", http.StatusBadRequest)
		return
	}

	// Return updated task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Delete task
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	taskID, err := utils.VarToUint(r, "id")

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Delete task
	deleteTaskResult := database.DB.Delete(&models.Task{}, taskID)

	if deleteTaskResult.Error != nil {
		http.Error(w, "Error while deleting task", http.StatusBadRequest)
		return
	}

	// Return success msg
	// Send a plain text response message
	message := fmt.Sprintf("Task with ID %d has been deleted", taskID)

	// Set response content type to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Write the string response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))

}
