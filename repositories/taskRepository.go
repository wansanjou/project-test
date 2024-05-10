package repositories

import (
	"github.com/wansanjou/project-test/modules"
	"gorm.io/gorm"
)


type TaskRepository struct {
	Client *gorm.DB
}

func NewTaskRepository(Client *gorm.DB) modules.TaskRepository {
	return &TaskRepository{Client: Client} 
}


func (t *TaskRepository) ListTasks(search, sortBy string) ([]modules.Task, error) {
	tasks := []modules.Task{}

	query := t.Client.Where("title LIKE ?", "%"+search+"%").
			Or("description LIKE ?", "%"+search+"%").
			Or("status LIKE ?", "%"+search+"%")

	switch sortBy {
	case "Title":
			query = query.Order("title ASC")
	case "Date":
			query = query.Order("date ASC")
	case "Status":
			query = query.Order("status ASC")
	}

	err := query.Find(&tasks).Error
	if err != nil {
			return nil, err
	}

	return tasks, nil
}


func (t *TaskRepository) GetTaskByID(id int) (*modules.Task , error) {
	tasks := modules.Task{}
	result := t.Client.First(&tasks , id)
	if result.Error != nil {
		return nil , result.Error
	}

	return &tasks , nil
}

func (t *TaskRepository) CreateTask(task modules.Task) (*modules.Task , error)  {
	err := t.Client.Create(&task).Error
	if err != nil {
		return nil, err
	}

	return &task , nil
}

func (t *TaskRepository) UpdateTask(id int , task modules.Task) (*modules.Task , error) {
	result := t.Client.Where("id = ?", id).Updates(task)
	if result.Error != nil {
			return nil, result.Error
	}

	return &task, nil
}


