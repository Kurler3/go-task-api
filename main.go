package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kurler3/go-task-api/database"
	"github.com/Kurler3/go-task-api/handlers"
	"github.com/Kurler3/go-task-api/utils"
	"github.com/gorilla/mux"
)

func main() {

	// Load env vars
	utils.LoadEnv()

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
	// Auth routes -------- //
	// -------------------- //

	// Register
	router.HandleFunc("/register", handlers.HandleRegister).Methods("POST")

	// Login
	router.HandleFunc("/login", handlers.HandleLogin).Methods("POST")

	// -------------------- //
	// User routes -------- //
	// -------------------- //

	// Get user by id

	// Update user

	// Delete user

	// Start the server
	port := "8081"

	fmt.Println("Server is running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
