package service

import (
	"gorm.io/gorm"
)

type StatusTask string

const (
	IN_PROGRESS StatusTask = "IN_PROGRESS"
	COMPLETED   StatusTask = "COMPLETED"
)

type TaskResponse struct {
	gorm.Model
	Title       string     `json:"title" validate:"required,max=100"`
	Description string     `json:"description"`
	Date       	string  		`json:"date" validate:"required"`
	Image       string     `json:"image"`
	Status      StatusTask `json:"status" validate:"required,oneof=IN_PROGRESS COMPLETED"`
}

type TaskService interface {
	ListTasks(search, sortBy string) ([]TaskResponse, error)
	GetTaskByID(id int) (*TaskResponse, error)
	CreateTask(task TaskResponse) (*TaskResponse, error)
	UpdateTask(id int, task TaskResponse) (*TaskResponse, error)
}
