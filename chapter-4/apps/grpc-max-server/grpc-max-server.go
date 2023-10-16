package grpcmaxserver

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	proto.UnimplementedMaxServiceServer
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	serverPort := 50053
	servAddr := fmt.Sprintf("0.0.0.0:%d", serverPort)

	fmt.Println("starting max gRPC server at", servAddr)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	proto.RegisterMaxServiceServer(s, a)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {

}

func (a *App) FindMax(stream proto.MaxService_FindMaxServer) error {
	max := int32(-1)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}

		num := req.GetNumber()
		if max < num {
			max = num
		}

		time.Sleep(2000 * time.Millisecond)
		err = stream.Send(&proto.MaxResponse{Max: max})
		if err != nil {
			log.Fatalf("Error while sending to client: %v", err)
			return err
		}
	}
}
