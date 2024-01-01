package grpcbooksserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/middleware"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/model"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
)

func (a *App) AddBook(_ context.Context, req *proto.Book) (*proto.AddBookResponse, error) {
	middleware.Incoming_api_req_counter.Add(1)
	log.Println("adding book")

	book := &model.DBBook{
		Isbn:      int(req.Isbn),
		Name:      req.Name,
		Publisher: req.Publisher,
	}

	a.bookRepo.AddBook(book)

	middleware.Book_create_pass_counter.Add(1)

	return &proto.AddBookResponse{Status: fmt.Sprintf("book with isbn(%d), name(%s), publisher(%s) added successfully", book.Isbn, book.Name, book.Publisher)}, nil
}

func (a *App) UpdateBook(_ context.Context, req *proto.Book) (*proto.UpdateBookResponse, error) {
	middleware.Incoming_api_req_counter.Add(1)
	log.Println("updating book")

	book := &model.DBBook{
		Isbn:      int(req.Isbn),
		Name:      req.Name,
		Publisher: req.Publisher,
	}

	a.bookRepo.UpdateBook(book)

	middleware.Book_update_pass_counter.Add(1)

	return &proto.UpdateBookResponse{Status: fmt.Sprintf("book with isbn(%d), name(%s), publisher(%s) updated successfully", book.Isbn, book.Name, book.Publisher)}, nil
}

func (a *App) ListBooks(ctx context.Context, _ *proto.Empty) (*proto.ListBooksRespose, error) {
	middleware.Incoming_api_req_counter.Add(1)
	log.Println("listing books")

	books, err := a.bookRepo.GetAllBooks()
	if err != nil {
		middleware.Book_get_fail_counter.Add(1)
		return nil, err
	}

	b, err := json.Marshal(books)
	if err != nil {
		middleware.Book_get_fail_counter.Add(1)
		return nil, fmt.Errorf("error while marshalling books", err.Error())
	}

	pbBooks := []*proto.Book{}
	err = json.Unmarshal(b, &pbBooks)
	if err != nil {
		middleware.Book_get_fail_counter.Add(1)
		return nil, fmt.Errorf("error while unmarshalling books", err.Error())
	}

	middleware.Book_get_pass_counter.Add(1)

	return &proto.ListBooksRespose{Books: pbBooks}, nil
}

func (a *App) GetBook(_ context.Context, req *proto.GetBookRequest) (*proto.Book, error) {
	middleware.Incoming_api_req_counter.Add(1)
	log.Println("fetching book")

	book := a.bookRepo.GetBook(int(req.Isbn))

	middleware.Book_get_pass_counter.Add(1)

	return &proto.Book{
		Isbn:      int32(book.Isbn),
		Name:      book.Name,
		Publisher: book.Publisher,
	}, nil
}

func (a *App) RemoveBook(_ context.Context, req *proto.RemoveBookRequest) (*proto.RemoveBookResponse, error) {
	middleware.Incoming_api_req_counter.Add(1)
	log.Println("removing book")

	a.bookRepo.RemoveBook(int(req.Isbn))

	middleware.Book_delete_pass_counter.Add(1)

	return &proto.RemoveBookResponse{Status: fmt.Sprintf("book with isbn(%d) removed successfully", req.Isbn)}, nil
}
