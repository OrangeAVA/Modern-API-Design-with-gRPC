package grpcbooksserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/model"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
)

func (a *App) AddBook(_ context.Context, req *proto.Book) (*proto.AddBookResponse, error) {
	log.Println("adding book")

	book := &model.DBBook{
		Isbn:      int(req.Isbn),
		Name:      req.Name,
		Publisher: req.Publisher,
	}

	a.bookRepo.AddBook(book)

	return &proto.AddBookResponse{Status: fmt.Sprintf("book with isbn(%d), name(%s), publisher(%s) added successfully", book.Isbn, book.Name, book.Publisher)}, nil
}

func (a *App) UpdateBook(_ context.Context, req *proto.Book) (*proto.UpdateBookResponse, error) {
	log.Println("updating book")

	book := &model.DBBook{
		Isbn:      int(req.Isbn),
		Name:      req.Name,
		Publisher: req.Publisher,
	}

	a.bookRepo.UpdateBook(book)

	return &proto.UpdateBookResponse{Status: fmt.Sprintf("book with isbn(%d), name(%s), publisher(%s) updated successfully", book.Isbn, book.Name, book.Publisher)}, nil
}

func (a *App) ListBooks(ctx context.Context, _ *proto.Empty) (*proto.ListBooksRespose, error) {
	log.Println("listing books")

	books, err := a.bookRepo.GetAllBooks()
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(books)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling books", err.Error())
	}

	pbBooks := []*proto.Book{}
	err = json.Unmarshal(b, &pbBooks)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling books", err.Error())
	}

	return &proto.ListBooksRespose{Books: pbBooks}, nil
}

func (a *App) GetBook(_ context.Context, req *proto.GetBookRequest) (*proto.Book, error) {
	log.Println("fetching book")

	book := a.bookRepo.GetBook(int(req.Isbn))

	return &proto.Book{
		Isbn:      int32(book.Isbn),
		Name:      book.Name,
		Publisher: book.Publisher,
	}, nil
}

func (a *App) RemoveBook(_ context.Context, req *proto.RemoveBookRequest) (*proto.RemoveBookResponse, error) {
	log.Println("removing book")

	a.bookRepo.RemoveBook(int(req.Isbn))

	return &proto.RemoveBookResponse{Status: fmt.Sprintf("book with isbn(%d) removed successfully", req.Isbn)}, nil
}
