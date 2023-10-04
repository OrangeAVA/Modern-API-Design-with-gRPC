package restbooksclient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/configs"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/model"
)

type App struct {
	serverAddr string
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	appConfig, err := configs.ProvideAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	a.serverAddr = appConfig.ClientConfig.ServerAddress

	a.performClientOperations()
}

func (a *App) Shutdown() {

}

func (a *App) performClientOperations() {
	book := model.DBBook{
		Isbn:      12348,
		Name:      "Sample Book",
		Publisher: "Sample Publisher",
	}

	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// add a book
	a.AddBook(jsonData)

	// update a book
	book.Name = "Sample Book Vol2"

	jsonData, err = json.Marshal(book)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	a.UpdateBook(jsonData)

	// list books
	a.ListBooks()

	// fetch book
	a.FetchBook(12345)

	// remove book
	a.RemoveBook(12348)
}
