package authserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/configs"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/jwt"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/middleware"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type App struct {
	proto.UnimplementedAuthServiceServer

	userStore  repo.UserStore
	jwtManager *jwt.JWTManager
}

func NewApp(userStore repo.UserStore, jwtManager *jwt.JWTManager) *App {
	return &App{userStore: userStore, jwtManager: jwtManager}
}

func (a *App) Start() {

	appConfig, err := configs.ProvideAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	servAddr := fmt.Sprintf("0.0.0.0:%d", appConfig.ServerConfig.Port)

	fmt.Println("starting auth gRPC server at", servAddr)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	middlewareOpts := middleware.ProvideGrpcMiddlewareServerOpts()
	opts = append(opts, middlewareOpts...)

	s := grpc.NewServer(opts...)

	proto.RegisterAuthServiceServer(s, a)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {}

func (a *App) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := a.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.InvalidArgument, "invlid username or password")
	}

	token, err := a.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token: %v", err)
	}

	res := &proto.LoginResponse{
		AccessToken: token,
	}

	return res, nil
}
