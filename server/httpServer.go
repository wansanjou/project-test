package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/wansanjou/project-test/apps/http/docs"
	"github.com/wansanjou/project-test/config"
	"github.com/wansanjou/project-test/handler"
	"gorm.io/gorm"
)

type IHttpServer interface {
	StartTaskServer(taskHandler *handler.TaskHandler)
}

type httpServer struct {
	app      *fiber.App
	cfg      *config.Config
	dbClient *gorm.DB
}

func NewHttpServer(cfg *config.Config, dbClient *gorm.DB) IHttpServer {
	return &httpServer{
		app:      fiber.New(),
		cfg:      cfg,
		dbClient: dbClient,
	}
}

func (s *httpServer) Listen() {
	log.Printf("Server is starting on %s", s.cfg.App.Url)
	if err := s.app.Listen(s.cfg.App.Url); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

// @title Task API
// @description This is a sample server for a task API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @in header
// @name Authorization
func (s *httpServer) StartTaskServer(taskHandler *handler.TaskHandler) {
	s.app.Get("/task", taskHandler.ListTasks)
	s.app.Get("/task/:id", taskHandler.GetTaskByID)
	s.app.Post("/task", taskHandler.CreateTask)
	s.app.Put("/task/:id", taskHandler.UpdateTask)
	s.app.Get("/swagger/*", swagger.HandlerDefault)

	log.Printf("Server is starting on %s", s.cfg.App.Url)
	if err := s.app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}


}