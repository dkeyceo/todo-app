package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/dkeyceo/todo-app"
	"github.com/dkeyceo/todo-app/pkg/handler"
	"github.com/dkeyceo/todo-app/pkg/repository"
	"github.com/dkeyceo/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization config: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		log.Fatalf("error loading env vars: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("error during init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
