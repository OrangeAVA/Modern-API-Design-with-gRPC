package main

import (
	bookserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-8/book-server"
)

func main() {
	app := bookserver.NewApp()
	app.Start()

	app.Shutdown()
}
