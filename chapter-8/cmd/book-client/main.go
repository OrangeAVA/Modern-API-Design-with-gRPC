package main

import (
	bookclient "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/book-client"
)

func main() {
	app := bookclient.NewApp("0.0.0.0:50091")
	app.Start()

	app.Shutdown()
}
