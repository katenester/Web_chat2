package main

import (
	"context"
	_ "github.com/jmoiron/sqlx"
	"github.com/katenester/Web_chat2/backend/internal/repository"
	"github.com/katenester/Web_chat2/backend/internal/repository/postgres/config"
	"github.com/katenester/Web_chat2/backend/internal/service"
	"github.com/katenester/Web_chat2/backend/internal/transport"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initalization config %s", err.Error())
	}

	// Иницализируем сервер
	srv := new(transport.Server)
	// Загрузка бд
	db, err := config.NewPostgresDB(config.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.dbpassword"),
	})
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
func initConfig() error {
	viper.AddConfigPath("backend/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
