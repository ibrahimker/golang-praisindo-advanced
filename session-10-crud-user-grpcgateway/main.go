package main

import (
	grpcHandler "github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/handler/grpc"
	pb "github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/proto/user_service/v1"
	"github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/repository/postgres_gorm"
	"github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/service"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	// setup gorm connection
	dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	// setup service

	// uncomment to use postgres gorm
	userRepo := postgres_gorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	//userHandler := ginHandler.NewUserHandler(userService)
	userHandler := grpcHandler.NewUserHandler(userService)

	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running grpc server in port :50051")
	_ = grpcServer.Serve(lis)
}
