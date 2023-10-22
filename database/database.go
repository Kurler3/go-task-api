package database

import (
	"github.com/Kurler3/go-task-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	// Create an SQLite connection with an in-memory database
	// Connect to the SQLite database
	db, err := gorm.Open("sqlite3", "tasks.db")

	// Check if err
	if err != nil {
		return nil, err
	}

	// Enable Gorm Logger, if needed
	db.LogMode(true)

	// Automigrate the Task model (create the tasks table)
	db.AutoMigrate(&models.Task{})
	// Migrate user model
	// db.AutoMigrate(&models.User{})

	// Assign the db instance to the package-level variable
	DB = db

	return db, nil

}
