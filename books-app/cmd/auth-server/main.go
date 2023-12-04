package main

import (
	authserver "github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/apps/auth-server"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/jwt"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/repo"
)

func main() {
	jwtMgr := jwt.NewJWTManager()
	userStore := repo.NewInMemoryUserStore()

	app := authserver.NewApp(userStore, jwtMgr)
	app.Start()

	app.Shutdown()
}
