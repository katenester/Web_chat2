package main

import (
	"context"
	"github.com/katenester/Web_chat2/backend/internal/repository"
	"github.com/katenester/Web_chat2/backend/internal/repository/sqllite/config"
	"github.com/katenester/Web_chat2/backend/internal/service"
	"github.com/katenester/Web_chat2/backend/internal/transport"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Иницализируем сервер
	srv := new(transport.Server)
	// Загрузка бд
	db, err := config.NewSQLLite()
	if err != nil {
		logrus.Fatal(err)
	}
	// Dependency injection for architecture application
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := transport.NewHandler(services)
	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server %s", err.Error())
		}
	}()

	logrus.Print("todo server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occured while shutting down server %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Fatalf("error occured while closing db %s", err.Error())
	}
}
