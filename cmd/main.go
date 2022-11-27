package main

import (
	"github.com/Dimiyrg/GoProject4"
	"github.com/Dimiyrg/GoProject4/pkg/handler"
	"github.com/Dimiyrg/GoProject4/pkg/repository"
	"github.com/Dimiyrg/GoProject4/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(GoProject4.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
