package main

import restfibonacciserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/apps/rest-fibonacci-server"

func main() {
	app := restfibonacciserver.NewApp()
	app.Start()
}
