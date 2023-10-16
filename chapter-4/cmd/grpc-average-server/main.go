package main

import grpcaverageserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/apps/grpc-average-server"

func main() {
	app := grpcaverageserver.NewApp()
	app.Start()
}
