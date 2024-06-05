package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-1-introduction-gin/handler"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	r.POST("/post", handler.PostHandler)
}
