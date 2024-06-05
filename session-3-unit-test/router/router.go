package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-3-unit-test/handler"
	"github.com/ibrahimker/golang-praisindo-advanced/session-3-unit-test/middleware"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	// Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndpoint := r.Group("/private")
	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		privateEndpoint.POST("/post", handler.PostHandler)
	}

}
