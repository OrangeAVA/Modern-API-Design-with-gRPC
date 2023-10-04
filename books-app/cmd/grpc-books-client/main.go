package main

import (
	grpcbooksclient "github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/apps/grpc-books-client"
)

func main() {
	app := grpcbooksclient.NewApp()
	app.Start()

	app.Shutdown()
}
