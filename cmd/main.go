package main

import (
	"github.com/p12s/wildberries-http-api/pkg/handler"
	"github.com/p12s/wildberries-http-api/pkg/repository"
	"github.com/p12s/wildberries-http-api/pkg/service"
	"github.com/p12s/wildberries-http-api"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

// TODO добавить доку
// TODO переделать логирование на Zap

func newHandler2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, man!\n\n")
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error init configs: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize DB: $s\n", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	/*http.HandleFunc("/lo2", newHandler2)
	srv := new(Server)
	if err := srv.Run(viper.GetString("port"), nil); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}*/

	srv := new(Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
