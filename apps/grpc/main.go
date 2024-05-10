package main

import (
	"os"

	"github.com/wansanjou/project-test/config"
	"github.com/wansanjou/project-test/pkg/db"
	"github.com/wansanjou/project-test/server"

)

func main() {
	cfg := config.NewConfig(func() string {
		if len(os.Args) > 1 {
			return os.Args[1]
		}
		return "./.env.grpc.task"
	}())

	dbClient := db.InitializeDB()

	grpcServer := server.NewGrpcServer(cfg, dbClient)	
	
	// taskRepository := repositories.NewTaskRepository(dbClient)
	// taskService := service.NewTaskService(taskRepository)
	// taskHandler := handler.NewTaskHandler(taskService)

	grpcServer.StartTaskServer()
}
