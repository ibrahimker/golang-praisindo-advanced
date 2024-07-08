package main

import (
	"context"
	"fmt"
	pb "github.com/ibrahimker/golang-praisindo-advanced/session-10-crud-user-grpcgateway/proto/user_service/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	runClient()
}

func runClient() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userServiceClient := pb.NewUserServiceClient(conn)

	ctx := context.Background()
	// create user
	resCreate, err := userServiceClient.CreateUser(ctx, &pb.CreateUserRequest{
		Name:     "test-user",
		Email:    "test-email@email.com",
		Password: "password",
	})
	fmt.Println(resCreate)

	// get all user
	resGetAll, err := userServiceClient.GetUsers(ctx, &emptypb.Empty{})
	fmt.Println(resGetAll)

	// delete user
	for _, u := range resGetAll.GetUsers() {
		resDel, _ := userServiceClient.DeleteUser(ctx, &pb.DeleteUserRequest{Id: u.GetId()})
		fmt.Println(resDel)
	}

}
