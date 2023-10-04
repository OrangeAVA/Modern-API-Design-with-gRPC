package restbooksserver

import (
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/model"
)

type Mapper interface {
	DBBook(*model.Book) *model.DBBook
	Book(*model.DBBook) *model.Book
}

func DBBook(book *model.Book) *model.DBBook {
	dbBook := &model.DBBook{
		Isbn:      book.Isbn,
		Name:      book.Name,
		Publisher: book.Publisher,
	}

	return dbBook
}

func Book(dbBooks *model.DBBook) *model.Book {
	book := &model.Book{
		Isbn:      dbBooks.Isbn,
		Name:      dbBooks.Name,
		Publisher: dbBooks.Publisher,
	}

	return book
}
