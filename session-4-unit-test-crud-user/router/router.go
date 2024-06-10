package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-unit-test-crud-user/handler"
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-unit-test-crud-user/middleware"
)

func SetupRouter(r *gin.Engine, userHandler *handler.UserHandler) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/:id", userHandler.GetUser)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)

	usersPrivateEndpoint := r.Group("/users")
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	usersPrivateEndpoint.POST("/", userHandler.CreateUser)
	usersPrivateEndpoint.PUT("/:id", userHandler.UpdateUser)
	usersPrivateEndpoint.DELETE("/:id", userHandler.DeleteUser)
}
