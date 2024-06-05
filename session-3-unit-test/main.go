package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-3-unit-test/router"
	"log"
)

func main() {
	// Inisialisasi router Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Definisikan route
	router.SetupRouter(r)

	// Jalankan server pada port 8080
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
