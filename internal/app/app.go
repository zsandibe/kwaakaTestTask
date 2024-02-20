package app

import (
	"context"
	"kwaaka-task/config"
	"kwaaka-task/internal/delivery"
	"kwaaka-task/internal/repository"
	"kwaaka-task/internal/server"
	"kwaaka-task/internal/service"
	"kwaaka-task/internal/storage"
	"kwaaka-task/pkg"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	config := config.NewConfig()
	pkg.InfoLog.Println("Config loaded")
	storage, err := storage.NewMongoDb(config)
	if err != nil {
		pkg.ErrorLog.Println(err)
		return
	}
	pkg.InfoLog.Println("Storage loaded")

	repository := repository.NewRepository(storage, config)
	pkg.InfoLog.Println("Repository loaded")

	service := service.NewService(repository, config)
	pkg.InfoLog.Println("Service loaded")

	delivery := delivery.NewHandler(service)
	pkg.InfoLog.Println("Delivery loaded")

	server := server.NewServer(config, delivery.Routes())

	go func() {
		if err := server.Run(); err != nil {
			pkg.ErrorLog.Printf("failed to start server: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		pkg.ErrorLog.Printf("failed to shutdown server: %v")
		return
	}

}
