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
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type App struct {
	proto.UnimplementedBookInfoServiceServer

	dbConn   *gorm.DB
	bookRepo *repo.BookRepository

	bookServerConn   *grpc.ClientConn
	bookServerClient proto.BookServiceClient

	reviewServerConn   *grpc.ClientConn
	reviewServerClient proto.ReviewServiceClient
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

	a.dialReviewServer(appConfig, err)
	a.dialBookServer(appConfig, err)

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

	proto.RegisterBookInfoServiceServer(s, a)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) dialReviewServer(appConfig *configs.AppConfig, err error) {
	reviewServAddr := appConfig.ServerConfig.ReviewServerAddress

	opts := grpc.WithInsecure()

	a.reviewServerConn, err = grpc.Dial(reviewServAddr, opts)
	if err != nil {
		log.Fatalf("could not connect review server: %v", err)
	}

	a.reviewServerClient = proto.NewReviewServiceClient(a.reviewServerConn)
}

func (a *App) dialBookServer(appConfig *configs.AppConfig, err error) {
	bookServAddr := appConfig.ServerConfig.BookServerAddress

	opts := grpc.WithInsecure()

	a.bookServerConn, err = grpc.Dial(bookServAddr, opts)
	if err != nil {
		log.Fatalf("could not connect books server: %v", err)
	}

	a.bookServerClient = proto.NewBookServiceClient(a.bookServerConn)
}

func (a *App) Shutdown() {
	dbInstance, _ := a.dbConn.DB()
	_ = dbInstance.Close()
}
