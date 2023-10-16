package main

import grpcaverageclient "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/apps/grpc-average-client"

func main() {
	app := grpcaverageclient.NewApp()
	app.Start()

	app.Shutdown()
}
