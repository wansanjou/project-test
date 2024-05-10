package main

import (
	"os"

	"github.com/wansanjou/project-test/config"
	"github.com/wansanjou/project-test/handler"

	"github.com/wansanjou/project-test/pkg/db"
	"github.com/wansanjou/project-test/repositories"
	"github.com/wansanjou/project-test/server"
	"github.com/wansanjou/project-test/service"
)

// @title Task API
// @description This is a sample server for a task API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @in header
// @name Authorization
func main() {
	// app := fiber.New()

	// db := db.InitializeDB()

	// taskRepository := repositories.NewTaskRepository(db)
	// taskService := service.NewTaskReposity(taskRepository)
	// taskHandler := handler.NewTaskHandler(taskService)

	// app.Get("/task",taskHandler.ListTasks)
	// app.Get("/task/:id",taskHandler.GetTaskByID)
	// app.Post("/task",taskHandler.CreateTask)
	// app.Put("/task/:id",taskHandler.UpdateTask)

	// app.Listen(":8080")

	// สร้าง config
	cfg := config.NewConfig(func() string {
		if len(os.Args) > 1 {
			return os.Args[1]
		}
		return "./.env.http.task"
	}())

	dbClient := db.InitializeDB()

	httpServer := server.NewHttpServer(cfg, dbClient)

	taskRepository := repositories.NewTaskRepository(dbClient)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)
	httpServer.StartTaskServer(taskHandler)

}