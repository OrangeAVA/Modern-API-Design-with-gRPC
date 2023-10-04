package restbooksserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/model"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/service"
	"github.com/gorilla/mux"
)

type BooksHandler struct {
	bookService service.BooksService
}

func GetNewBooksHandler(bookService service.BooksService) *BooksHandler {
	return &BooksHandler{bookService: bookService}
}

func (bh *BooksHandler) BooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookService.GetAllBooks()
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, books)
}

func (bh *BooksHandler) AddOrRemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	muxVar := mux.Vars(r)
	isbnStr := muxVar["isbn"]
	isbn, err := strconv.Atoi(isbnStr)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if r.Method == "GET" {
		book, err := bh.bookService.GetBook(isbn)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, book)
	}

	if r.Method == "DELETE" {
		bh.bookService.RemoveBook(isbn)
		respondWithJSON(w, http.StatusOK, map[string]string{SuccessResponseFieldKey: "book removed"})
	}
}

func (bh *BooksHandler) UpsertBookHandler(w http.ResponseWriter, r *http.Request) {
	var book *model.Book
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if r.Method == "POST" {
		bh.bookService.AddBook(DBBook(book))
	} else if r.Method == "PUT" {
		bh.bookService.UpdateBook(DBBook(book))
	} else {
		respondWithError(w, http.StatusMethodNotAllowed, "invalid request method")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{SuccessResponseFieldKey: "book upserted"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{ErrorResponseFieldKey: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
