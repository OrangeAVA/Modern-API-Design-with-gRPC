package restbooksserver

import (
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/service"
	"github.com/gorilla/mux"
)

func ProvideRouter(bookService service.BooksService) *mux.Router {
	r := mux.NewRouter()

	booksHandler := GetNewBooksHandler(bookService)

	r.HandleFunc("/books/all", booksHandler.BooksHandler).Methods("GET")
	r.HandleFunc("/books", booksHandler.UpsertBookHandler).Methods("POST", "PUT")
	r.HandleFunc("/books/{isbn:[0-9]+}", booksHandler.AddOrRemoveBookHandler).Methods("GET", "DELETE")

	return r
}
