package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/handler"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/repository/postgres_gorm"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/router"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup gorm connection
	dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	// setup repository
	userRepo := postgres_gorm.NewUserRepository(gormDB)
	submissionRepo := postgres_gorm.NewSubmissionRepository(gormDB)

	// service and handler declaration
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	submissionService := service.NewSubmissionService(submissionRepo)
	submissionHandler := handler.NewSubmissionHandler(submissionService)

	// Routes
	router.SetupRouter(r, userHandler, submissionHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
