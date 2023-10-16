package main

import restfibonacciclient "github.com/HiteshRepo/Modern-API-Design-with-gRPC/chapter-4/apps/rest-fibonacci-client"

func main() {
	app := restfibonacciclient.NewApp()
	app.Start()
}
