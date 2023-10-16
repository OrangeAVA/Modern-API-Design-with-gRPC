package main

import grpcmaxserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/apps/grpc-max-server"

func main() {
	app := grpcmaxserver.NewApp()
	app.Start()
}
