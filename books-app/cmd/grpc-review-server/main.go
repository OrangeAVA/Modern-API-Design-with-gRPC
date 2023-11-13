package main

import (
	grpcreviewserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/apps/grpc-review-server"
)

func main() {
	app := grpcreviewserver.NewApp()
	app.Start()

	app.Shutdown()
}
