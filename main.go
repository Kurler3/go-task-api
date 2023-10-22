package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kurler3/go-task-api/database"
	"github.com/Kurler3/go-task-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	_, err := database.InitDB()
	if err != nil {
		// Handle error
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	// Create a new Mux router
	router := mux.NewRouter()

	////////////////////////////////////////
	// Define API routes ///////////////////
	////////////////////////////////////////

	// -------------------- //
	// Task routes -------- //
	// -------------------- //

	// Create task
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")

	// Get all tasks (for userId from jwt)
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")

	// Get task by id
	router.HandleFunc("/tasks/{id}", handlers.GetTaskById).Methods("GET")

	// Update task
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PATCH")

	// Delete task
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// -------------------- //
	// User routes -------- //
	// -------------------- //

	// Register

	// Login

	// Get user by id

	// Update user

	// Delete user

	// Start the server
	port := "8081"

	fmt.Println("Server is running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
