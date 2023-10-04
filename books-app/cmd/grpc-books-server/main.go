package main

import (
	grpcbooksserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/apps/grpc-books-server"
)

func main() {
	app := grpcbooksserver.NewApp()
	app.Start()

	app.Shutdown()
}
