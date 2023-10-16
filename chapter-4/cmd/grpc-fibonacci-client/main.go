package main

import grpcfibonacciclient "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/apps/grpc-fibonacci-client"

func main() {
	app := grpcfibonacciclient.NewApp()
	app.Start()

	app.Shutdown()
}
