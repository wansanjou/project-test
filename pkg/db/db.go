package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wansanjou/project-test/modules"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeDB() *gorm.DB {
	const (
		host     = "localhost"  // or the Docker service name if running in another container
		port     = 5432         // default PostgreSQL port
		user     = "myuser"     // as defined in docker-compose.yml
		password = "mypassword" // as defined in docker-compose.yml
		dbname   = "mydatabase" // as defined in docker-compose.yml
	)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger
	})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&modules.Task{})
	return db
}
