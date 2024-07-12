package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcHandler "github.com/ibrahimker/golang-praisindo-advanced/session-11-crud-user-grpcgateway-cache/handler/grpc"
	pb "github.com/ibrahimker/golang-praisindo-advanced/session-11-crud-user-grpcgateway-cache/proto/user_service/v1"
	"github.com/ibrahimker/golang-praisindo-advanced/session-11-crud-user-grpcgateway-cache/repository/postgres_gorm_raw"
	"github.com/ibrahimker/golang-praisindo-advanced/session-11-crud-user-grpcgateway-cache/service"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"time"
)

func main() {
	// setup gorm connection
	dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// setup service

	// uncomment to use postgres gorm
	userRepo := postgres_gorm_raw.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo, rdb)
	userHandler := grpcHandler.NewUserHandler(userService)

	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Println("Running grpc server in port :50051")
		_ = grpcServer.Serve(lis)
	}()
	time.Sleep(1 * time.Second)

	// Run the grpc gateway
	conn, err := grpc.NewClient(
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	if err = pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// dengan default http server
	/*
		gwServer := &http.Server{
			Addr:    ":8080",
			Handler: gwmux,
		}
		log.Println("Running grpc gateway server in port :8080")
		_ = gwServer.ListenAndServe()
	*/

	// dengan GIN
	gwServer := gin.Default()
	gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))
	log.Println("Running grpc gateway server in port :8080")
	_ = gwServer.Run(":8080")
}
