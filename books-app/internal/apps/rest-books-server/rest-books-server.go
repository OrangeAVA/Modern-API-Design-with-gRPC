package restbooksserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/configs"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/db"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/db/migrations"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/repo"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/service"
	"gorm.io/gorm"
)

type App struct {
	dbConn *gorm.DB
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

	migrator, err := migrations.ProvideMigrator(appConfig.DBConfig, dbConn)
	if err != nil {
		log.Fatal(err)
	}

	migrator.RunMigrations()

	bookRepo := repo.GetNewBookRepository(dbConn)
	bookSrv := service.GetNewBooksService(bookRepo)
	r := ProvideRouter(bookSrv)

	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", appConfig.ServerConfig.Host, appConfig.ServerConfig.Port),
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server")
	log.Fatal(srv.ListenAndServe())
}

func (a *App) Shutdown() {
	dbInstance, _ := a.dbConn.DB()
	_ = dbInstance.Close()
}
