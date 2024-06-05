package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-2-latihan-crud-user-gin/handler"
	"github.com/ibrahimker/golang-praisindo-advanced/session-2-latihan-crud-user-gin/middleware"
)

func SetupRouter(r *gin.Engine) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/:id", handler.GetUser)
	usersPublicEndpoint.GET("/", handler.GetAllUsers)

	usersPrivateEndpoint := r.Group("/users")
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	usersPrivateEndpoint.POST("/", handler.CreateUser)
	usersPrivateEndpoint.PUT("/:id", handler.UpdateUser)
	usersPrivateEndpoint.DELETE("/:id", handler.DeleteUser)
}
