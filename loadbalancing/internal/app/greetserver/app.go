package greetserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/loadbalancing/internal/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type App struct {
	proto.UnimplementedGreetServiceServer
}

func (a *App) Start() {
	fmt.Println("Starting greet server")

	port := "50051"
	if len(os.Getenv("GRPC_PORT")) > 0 {
		port = os.Getenv("GRPC_PORT")
	}

	servAddr := fmt.Sprintf("0.0.0.0:%s", port)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: time.Minute * 5,
		}),
	}
	s := grpc.NewServer(opts...)
	proto.RegisterGreetServiceServer(s, &App{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {

}

func (a *App) Greet(_ context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	var podIp string
	var podName string
	if len(os.Getenv("POD_IP")) > 0 {
		podIp = os.Getenv("POD_IP")
	}

	if len(os.Getenv("POD_NAME")) > 0 {
		podName = os.Getenv("POD_NAME")
	}

	result := fmt.Sprintf("Hello, %s %s from pod: name(%s), ip(%s).", firstName, lastName, podName, podIp)
	return &proto.GreetingResponse{
		Result: result,
	}, nil
}
