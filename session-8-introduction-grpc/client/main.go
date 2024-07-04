package main

import (
	"context"
	pb "github.com/ibrahimker/golang-praisindo-advanced/session-8-introduction-grpc/proto/helloworld/v1"
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
	greeterClient := pb.NewGreeterServiceClient(conn)

	name := "world"
	r, err := greeterClient.SayHello(context.Background(), &pb.SayHelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
