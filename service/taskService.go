package service

import (
	"errors"
	"time"

	"github.com/wansanjou/project-test/modules"
)

type taskService struct {
	task_repo modules.TaskRepository
}

func NewTaskService(task_repo modules.TaskRepository) TaskService {
	return taskService{task_repo: task_repo}
}

func (s taskService) ListTasks(search, sortBy string) ([]TaskResponse, error) {
	tasks, err := s.task_repo.ListTasks(search, sortBy)
	if err != nil {
			return nil, err
	}

	task_responses := make([]TaskResponse, 0, len(tasks))
	for _, task := range tasks {
			task_response := TaskResponse{
					Title:       task.Title,
					Description: task.Description,
					Date:        task.Date,
					Image:       task.Image,
					Status:      StatusTask(task.Status),
			}
			task_responses = append(task_responses, task_response)
	}

	return task_responses, nil
}




func (s taskService) GetTaskByID(id int) (*TaskResponse, error)  {
	task , err := s.task_repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task_response := TaskResponse{
		Title: task.Title,
		Description: task.Description,
		Date: task.Date,
		Image: task.Image,
		Status: StatusTask(task.Status),
	}

	return &task_response , nil
}

func (s taskService) CreateTask(task_res TaskResponse) (*TaskResponse, error) {
	if task_res.Status == "" {
			task_res.Status = "IN_PROGRESS"
	} else if task_res.Status != "IN_PROGRESS" && task_res.Status != "COMPLETED" {
			return nil, errors.New("Error data Status !!")
	}

	if task_res.Date == "" {
		task_res.Date = time.Now().Format(time.RFC3339)
	}

	task := modules.Task{
			Title:       task_res.Title,
			Description: task_res.Description,
			Date:        task_res.Date, 
			Image:       task_res.Image,
			Status:      modules.StatusTask(task_res.Status),
	}

	insert_task, err := s.task_repo.CreateTask(task)
	if err != nil {
			return nil, err
	}

	response := TaskResponse{
			Title:       insert_task.Title,
			Description: insert_task.Description,
			Date:        task_res.Date,
			Image:       insert_task.Image,
			Status:      StatusTask(insert_task.Status),
	}

	return &response, nil
}



func (s taskService) UpdateTask(id int, task_res TaskResponse) (*TaskResponse, error) {
	if id == 0 {
		return nil, errors.New("Invalid ID")
	}

	if task_res.Status != "IN_PROGRESS" && task_res.Status != "COMPLETED" {
		return nil, errors.New("Error data Status !!")
	}

	task := modules.Task{
		Title:       task_res.Title,
		Description: task_res.Description,
		Date:        task_res.Date,
		Image:       task_res.Image,
		Status:      modules.StatusTask(task_res.Status),
	}

	update_task, err := s.task_repo.UpdateTask(id, task)
	if err != nil {
		return nil, errors.New("Failed to update task")
	}

	response := TaskResponse{
		Title:       update_task.Title,
		Description: update_task.Description,
		Date:        update_task.Date,
		Image:       update_task.Image,
		Status:      StatusTask(update_task.Status),
	}

	return &response, nil
}
