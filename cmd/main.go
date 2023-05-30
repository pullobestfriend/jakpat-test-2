package main

import (
	jakpatTest "jakpat-test-2"
	"jakpat-test-2/pkg/handler"
	"jakpat-test-2/pkg/repository"
	"jakpat-test-2/pkg/service"
	"jakpat-test-2/pkg/usecase"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error initializing configs: ", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env variables: ", err.Error())
	}
	//gin.SetMode(gin.ReleaseMode)
	//fmt.Println("Starting server...")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatal("error initializing database: ", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	usecases := usecase.NewUsecase(
		services,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl"),
	)
	handlers := handler.NewHandler(usecases)

	srv := new(jakpatTest.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occurred while running http server: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
