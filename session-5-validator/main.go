package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/handler"
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/repository/slice"
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/router"
	"github.com/ibrahimker/golang-praisindo-advanced/session-5-validator/service"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
