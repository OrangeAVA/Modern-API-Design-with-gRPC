package grpcfibonacciserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	proto.UnimplementedFibonacciServiceServer
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	serverPort := 50051
	servAddr := fmt.Sprintf("0.0.0.0:%d", serverPort)

	fmt.Println("starting fibonacci gRPC server at", servAddr)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	proto.RegisterFibonacciServiceServer(s, a)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {}

func (a *App) SyncFibonacci(ctx context.Context, req *proto.FibonacciRequest) (*proto.SyncFibonacciResponse, error) {
	numFibonacci := int(req.GetNumber())

	fibonacciNumbers := make([]int32, numFibonacci)

	now := time.Now()
	for i := 0; i < numFibonacci; i++ {
		fibonacciNumbers[i] = int32(fib(i))
	}
	timeTaken := time.Since(now).Seconds()

	resp := proto.SyncFibonacciResponse{
		TimeTaken:        fmt.Sprintf("%v seconds", timeTaken),
		FibonacciNumbers: fibonacciNumbers,
	}

	return &resp, nil
}

func (a *App) AsyncFibonacci(req *proto.FibonacciRequest, stream proto.FibonacciService_AsyncFibonacciServer) error {
	numFibonacci := int(req.GetNumber())

	for i := 0; i < numFibonacci; i++ {
		fibI := int32(fib(i))
		resp := &proto.AsyncFibonacciResponse{
			Sequence:        int32(i),
			FibonacciNumber: fibI,
		}
		stream.Send(resp)
	}

	return nil
}

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
