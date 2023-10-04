package main

import (
	restbooksclient "github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/apps/rest-books-client"
)

func main() {
	app := restbooksclient.NewApp()
	app.Start()

	app.Shutdown()
}
