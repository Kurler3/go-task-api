package models

type Task struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	UserID      uint   `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"dueDate"`
	// User        User   `json:"user"`
}
