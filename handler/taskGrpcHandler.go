package handler

import (
	"fmt"
	"io"
	"strconv"

	pbTask "github.com/wansanjou/project-test/proto/task"
	"github.com/wansanjou/project-test/repositories"
)

type TaskGrpcHandler struct {
	TaskRepository *repositories.TaskRepository
	pbTask.UnimplementedTaskServiceServer
}

func (h *TaskGrpcHandler) ListTasks(stream pbTask.TaskService_ListTasksServer) error {	
	tasks := &pbTask.TaskArr{
		Data: make([]*pbTask.Task, 0),
	}

	for {
			taskID, err := stream.Recv()
			if err == io.EOF {
					fmt.Println("task_id out of range")
					break
			}
			if err != nil {
					return err
			}

			id, err := strconv.Atoi(taskID.Id)
			if err != nil {
					return err
			}

			task, err := h.TaskRepository.GetTaskByID(id)
			if err != nil {
					return err
			}

			tasks.Data = append(tasks.Data, &pbTask.Task{
				Title: task.Title,
				Description: task.Description,
				Date: task.Date,
			})
	}

	stream.SendAndClose(tasks)
	return nil
}
