package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/loadbalancing/internal/app/greetserver"
)

func main() {
	a := greetserver.App{}
	a.Start()
	<-interrupt()
	a.Shutdown()
}

func interrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	return interrupt
}
