package main

import (
	"context"
	"fmt"
	pb "github.com/ibrahimker/golang-praisindo-advanced/session-8-introduction-grpc/proto/helloworld/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	hello := &pb.Hello{
		Id:        0,
		Name:      "test nama",
		Active:    false,
		Type:      pb.HelloType_HELLO_TYPE_ACTIVE,
		Schools:   []string{"sd 1", "smp 2"},
		CreatedAt: timestamppb.New(time.Now().Add(-4 * time.Hour)), // 4 hours ago
		UpdatedAt: timestamppb.New(time.Now()),
	}
	fmt.Println(hello)
	return &pb.SayHelloResponse{Message: fmt.Sprintf("Hello World %s", in.GetName()), Hello: hello}, nil
}

func main() {
	runServer()
}

func runServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, &server{})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
