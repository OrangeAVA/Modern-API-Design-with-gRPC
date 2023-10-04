package main

import (
	restbooksserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/apps/rest-books-server"
)

func main() {
	app := restbooksserver.NewApp()
	app.Start()

	app.Shutdown()
}
