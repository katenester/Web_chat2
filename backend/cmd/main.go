package main

import "github.com/katenester/Web_chat2/backend/internal/transport"

func main() {
	// Иницализируем сервер
	srv := new(transport.Server)
	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server %s", err.Error())
		}
	}()
}
