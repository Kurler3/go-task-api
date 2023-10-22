package models

type User struct {
	ID                uint   `json:"id" gorm:"primary_key"`
	Name              string `json:"name"`
	Email             string `json:"email" gorm:"uniqueIndex;not null"`
	EncryptedPassword string `json:"encryptedPassword"`
	Tasks             []Task `json:"tasks" gorm:"foreignkey:UserID"` // Define the relationship with tasks
}
