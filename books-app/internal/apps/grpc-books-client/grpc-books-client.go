package grpcbooksclient

import (
	"fmt"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/configs"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	serverConn *grpc.ClientConn
	client     proto.BookServiceClient
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	appConfig, err := configs.ProvideAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	servAddr := appConfig.ClientConfig.ServerAddress

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	a.serverConn, err = grpc.Dial(servAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	a.client = proto.NewBookServiceClient(a.serverConn)

	a.performClientOperations()
}

func (a *App) Shutdown() {
	a.serverConn.Close()
}

func (a *App) performClientOperations() {
	book := &proto.Book{
		Isbn:      12348,
		Name:      "atomic habits",
		Publisher: "random house business books",
	}

	// add a book
	a.AddBook(book)

	// list books
	found := false
	books := a.ListBooks()
	for i, b := range books {
		fmt.Printf("[%d.] Name(%s), Isbn(%d), Publisher(%s)", i, b.Name, b.Isbn, b.Publisher)
		if b.Isbn == book.Isbn &&
			b.Name == book.Name &&
			b.Publisher == book.Publisher {
			found = true
		}
	}
	if found {
		fmt.Println("Book sent through 'AddBook' request was verified to have been added while listing books.")
	}

	// fetch a book
	fetchedBook := a.FetchBook(12345)
	if fetchedBook.Isbn == 12345 &&
		fetchedBook.Name == "" &&
		fetchedBook.Publisher == "" {
		fmt.Println("Book added via migrations was successfully fetched.")
	}

	// update a book
	book.Name = "atomic habits vol-2"
	a.UpdateBook(book)

	// delete a book
	a.RemoveBook(int(book.Isbn))
}
