package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-unit-test-crud-user/entity"
	"github.com/ibrahimker/golang-praisindo-advanced/session-6-db-pgx/repository/postgres_pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//r := gin.Default()

	// "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"
	dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connect to db")
	var u entity.User
	ctx := context.Background()
	err = pool.QueryRow(ctx, "select id,name from users order by id desc limit 1").Scan(&u.ID, &u.Name)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("user retrieved", u)
	_, err = pool.Exec(ctx, "insert into users (name,email,password,created_at,updated_at) values "+
		"('test','test@test.com','pass',NOW(),NOW())")
	if err != nil {
		log.Fatalln(err)
	}
	err = pool.QueryRow(ctx, "select id,name from users order by id desc limit 1").Scan(&u.ID, &u.Name)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("user retrieved", u)
	// setup service

	// slice db is disabled. uncomment to enabled
	//var mockUserDBInSlice []entity.User
	//_ = slice.NewUserRepository(mockUserDBInSlice)

	// pgx db is enabled. comment to disabled
	//userRepo := postgres_pgx.NewUserRepository(pgxPool)
	//userService := service.NewUserService(userRepo)
	//userHandler := handler.NewUserHandler(userService)
	//
	//// Routes
	//router.SetupRouter(r, userHandler)
	//
	//// Run the server
	//log.Println("Running server on port 8080")
	//r.Run(":8080")
}

func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
	return pgxpool.New(context.Background(), dbURL)
}
