package modules

import (
	"gorm.io/gorm"
)

type StatusTask string

const (
	IN_PROGRESS StatusTask = "IN_PROGRESS"
	COMPLETED   StatusTask = "COMPLETED"
)

type Task struct {
	gorm.Model
	Title       string     `json:"title" validate:"required,max=100"`
	Description string     `json:"description"`
	Date        string  `json:"date" validate:"required"`
	Image       string     `json:"image"`
	Status      StatusTask `json:"status" validate:"required,oneof=IN_PROGRESS COMPLETED"`
}

type TaskRepository interface {
	ListTasks(search, sortBy string) ([]Task, error)
	GetTaskByID(id int) (*Task, error)
	CreateTask(task Task) (*Task, error)
	UpdateTask(id int, task Task) (*Task, error)
}
