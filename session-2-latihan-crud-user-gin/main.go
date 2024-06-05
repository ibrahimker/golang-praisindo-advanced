package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-2-latihan-crud-user-gin/router"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Routes
	router.SetupRouter(r)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
