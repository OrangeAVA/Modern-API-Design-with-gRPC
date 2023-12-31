package bookserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	proto.UnimplementedBookServiceServer
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	servAddr := "0.0.0.0:50091"

	fmt.Println("starting books gRPC server at", servAddr)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	proto.RegisterBookServiceServer(s, a)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {

}

func (a *App) ListBooks(ctx context.Context, _ *proto.Empty) (*proto.ListBooksRespose, error) {
	log.Println("listing books")

	books := a.GetAllBooks()

	return &proto.ListBooksRespose{Books: books}, nil
}

func (a *App) GetAllBooks() []*proto.Book {
	log.Println("listing books")

	pbBooks := []*proto.Book{
		{
			Isbn:      1234,
			Name:      "book-1",
			Publisher: "publisher-1",
		},
		{
			Isbn:      1235,
			Name:      "book-2",
			Publisher: "publisher-2",
		},
	}

	return pbBooks
}
