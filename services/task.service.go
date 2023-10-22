package services

import (
	"github.com/Kurler3/go-task-api/database"
	"github.com/Kurler3/go-task-api/models"
)

func CreateTask(task *models.Task) (*models.Task, error) {
	result := database.DB.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func UpdateTask(task *models.Task) (*models.Task, error) {
	result := database.DB.Save(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func DeleteTask(taskID uint) error {
	result := database.DB.Delete(&models.Task{}, taskID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTaskById(taskID uint) (*models.Task, error) {
	task := &models.Task{}
	result := database.DB.First(task, taskID)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}
