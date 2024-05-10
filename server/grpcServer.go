package server

import (
	"log"
	"net"

	"github.com/wansanjou/project-test/config"
	"github.com/wansanjou/project-test/handler"
	pbTask "github.com/wansanjou/project-test/proto/task"
	"github.com/wansanjou/project-test/repositories"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type grpcServer struct {
	cfg      *config.Config
	dbClient *gorm.DB
}

func NewGrpcServer(cfg *config.Config , dbClient *gorm.DB) *grpcServer {
	return &grpcServer{
		cfg: cfg,
		dbClient: dbClient ,
	}
}

func (s *grpcServer) StartTaskServer()  {
	opts := make([]grpc.ServerOption,0)
	gs := grpc.NewServer(opts...)

	log.Printf("%s Server is starting on %s " ,s.cfg.App.AppName, s.cfg.App.Url)
	lis ,err := net.Listen("tcp" , s.cfg.App.Url)
	if err != nil {
		log.Fatal(err)
	}

	taskHandler := &handler.TaskGrpcHandler{
    TaskRepository: &repositories.TaskRepository{
        Client: s.dbClient,
    	},
	}	

	pbTask.RegisterTaskServiceServer(gs, taskHandler)


	gs.Serve(lis)
}