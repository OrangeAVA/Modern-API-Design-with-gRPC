package grpcfibonacciclient

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/proto"
	"google.golang.org/grpc"
)

var (
	typeOfCall string
	number     int
)

func init() {
	flag.StringVar(&typeOfCall, "typeOfCall", "sync", "do you want to make sync or async calls to server?")
	flag.IntVar(&number, "number", 10, "for what number do you want the fibonacci sequence")
}

type App struct {
	clientConn *grpc.ClientConn
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	var err error
	flag.Parse()

	servAddr := "localhost:50051"

	opts := grpc.WithInsecure()
	a.clientConn, err = grpc.Dial(servAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	fibonacciServiceClient := proto.NewFibonacciServiceClient(a.clientConn)

	if typeOfCall == "sync" {
		syncFibonacci(fibonacciServiceClient, number)
	}

	if typeOfCall == "async" {
		asyncFibonacci(fibonacciServiceClient, number)
	}
}

func (a *App) Shutdown() {
	a.clientConn.Close()
}

func syncFibonacci(client proto.FibonacciServiceClient, number int) {
	req := &proto.FibonacciRequest{Number: int32(number)}

	res, err := client.SyncFibonacci(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling fibonacci rpc : %v", err)
	}

	fibResp := ""
	for _, n := range res.FibonacciNumbers {
		fibResp = fmt.Sprintf("%s,%d", fibResp, n)
	}

	log.Printf("Synchronous Fibonacci sequence result for %d: %v, and time taken was: %s\n", number, fibResp[1:], res.TimeTaken)
}

func asyncFibonacci(client proto.FibonacciServiceClient, number int) {
	req := &proto.FibonacciRequest{Number: int32(number)}

	resStream, err := client.AsyncFibonacci(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling fibonacci stream rpc : %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream from fibonacci rpc : %v", err)
		}
		log.Printf("Response form fibonacci stream: sequence(%d), fibonacci number(%d)", msg.Sequence, msg.FibonacciNumber)
	}
}
