package grpcbooksclient

import (
	"context"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
)

func (a *App) AddBook(book *proto.Book) {
	resp, err := a.client.AddBook(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Status)
}

func (a *App) ListBooks() []*proto.Book {
	resp, err := a.client.ListBooks(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	return resp.Books
}

func (a *App) FetchBook(isbn int) *proto.Book {
	resp, err := a.client.GetBook(context.Background(), &proto.GetBookRequest{Isbn: int32(isbn)})
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func (a *App) RemoveBook(isbn int) {
	resp, err := a.client.RemoveBook(context.Background(), &proto.RemoveBookRequest{Isbn: int32(isbn)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Status)
}

func (a *App) UpdateBook(book *proto.Book) {
	resp, err := a.client.UpdateBook(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Status)
}
