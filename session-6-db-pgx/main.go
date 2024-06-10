package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/handler"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/repository/postgres_pgx"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/repository/slice"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/router"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	pgxPool, err := connectDB("postgresql://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatalln(err)
	}
	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	_ = postgres_pgx.NewUserRepository(pgxPool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}

func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
	return pgxpool.New(context.Background(), dbURL)
}
