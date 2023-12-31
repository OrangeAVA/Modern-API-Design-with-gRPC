package bookclient

import (
	"context"
	"fmt"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/proto"
	"google.golang.org/grpc"
)

type App struct {
	serverAddr string

	serverConn *grpc.ClientConn
	client     proto.BookServiceClient
}

func NewApp(serverAddr string) *App {
	return &App{
		serverAddr: serverAddr,
	}
}

func (a *App) Start() {
	var err error

	opts := grpc.WithInsecure()

	a.serverConn, err = grpc.Dial(a.serverAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	a.client = proto.NewBookServiceClient(a.serverConn)

	a.ListBooks(context.Background(), &proto.Empty{})
}

func (a *App) Shutdown() {
	a.serverConn.Close()
}

func (a *App) ListBooks(ctx context.Context, in *proto.Empty, _ ...grpc.CallOption) (*proto.ListBooksRespose, error) {
	books, err := a.client.ListBooks(ctx, in)
	if err != nil {
		return nil, err
	}

	for _, b := range books.GetBooks() {
		fmt.Printf("----------> Isbn: %d, Name: %s, Publisher: %s\n", b.GetIsbn(), b.GetName(), b.GetPublisher())
	}

	return books, nil
}
