package main

import grpcmaxclient "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/apps/grpc-max-client"

func main() {
	app := grpcmaxclient.NewApp()
	app.Start()

	app.Shutdown()
}
