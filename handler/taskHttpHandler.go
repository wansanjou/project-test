package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wansanjou/project-test/service"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// @Summary Get all tasks
// @Description Get details of all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} Task
// @Router /task [get]
func (h *TaskHandler) ListTasks(c *fiber.Ctx) error {
	search := c.Query("search")
	sortBy := c.Query("sortBy")

	tasks, err := h.taskService.ListTasks(search,sortBy)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get tasks"})
	}

	return c.Status(http.StatusOK).JSON(tasks)
}

// @Summary Get task by ID
// @Description Get task details by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} Task
// @Failure 400 {object} ErrorResponse
// @Router /task/{id} [get]
func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	task, err := h.taskService.GetTaskByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get task"})
	}

	return c.JSON(task)
}

// @Summary Create a task
// @Description Create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body Task true "Task object"
// @Success 201 {object} Task
// @Router /task [post]
func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var task service.TaskResponse
	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := h.taskService.CreateTask(task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.Status(http.StatusCreated).JSON(response)
}

// @Summary Update a task
// @Description Update task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body Task true "Updated task object"
// @Success 200 {object} Task
// @Router /task/{id} [put]
func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var task service.TaskResponse
	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := h.taskService.UpdateTask(id, task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}

	return c.JSON(response)
}