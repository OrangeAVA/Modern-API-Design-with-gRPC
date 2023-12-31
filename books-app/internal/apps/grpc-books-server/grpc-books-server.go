package grpcbooksserver

import (
	"fmt"
	"log"
	"net"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/configs"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/db"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/db/migrations"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/middleware"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/repo"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type App struct {
	proto.UnimplementedBookServiceServer

	dbConn   *gorm.DB
	bookRepo *repo.BookRepository
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {

	appConfig, err := configs.ProvideAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.ProvideDBConn(&appConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	a.dbConn = dbConn

	a.bookRepo = repo.GetNewBookRepository(a.dbConn)

	migrator, err := migrations.ProvideMigrator(appConfig.DBConfig, dbConn)
	if err != nil {
		log.Fatal(err)
	}

	migrator.RunMigrations()

	servAddr := fmt.Sprintf("0.0.0.0:%d", appConfig.ServerConfig.Port)

	fmt.Println("starting books gRPC server at", servAddr)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	middlewareOpts := middleware.ProvideGrpcMiddlewareServerOpts()
	opts = append(opts, middlewareOpts...)

	s := grpc.NewServer(opts...)

	proto.RegisterBookServiceServer(s, a)

	reflection.Register(s)

	grpc_prometheus.Register(s)
	middleware.RunPrometheusServer(appConfig.ServerConfig.PrometheusPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {
	dbInstance, _ := a.dbConn.DB()
	_ = dbInstance.Close()
}
