package main

import grpcfibonacciserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/apps/grpc-fibonacci-server"

func main() {
	app := grpcfibonacciserver.NewApp()
	app.Start()
}
