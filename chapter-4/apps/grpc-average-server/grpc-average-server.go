package grpcaverageserver

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	proto.UnimplementedAverageServiceServer
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	serverPort := 50052
	servAddr := fmt.Sprintf("0.0.0.0:%d", serverPort)

	fmt.Println("starting average gRPC server at", servAddr)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	proto.RegisterAverageServiceServer(s, a)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {}

func (a *App) FindAverage(stream proto.AverageService_FindAverageServer) error {
	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.AverageResponse{
				Average: int32(sum / count),
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}

		count += 1
		sum += int(req.GetNumber())
	}
}
