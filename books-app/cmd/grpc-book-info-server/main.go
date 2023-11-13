package main

import (
	grpcbookinfoserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/apps/grpc-book-info-server"
)

func main() {
	app := grpcbookinfoserver.NewApp()
	app.Start()

	app.Shutdown()
}
