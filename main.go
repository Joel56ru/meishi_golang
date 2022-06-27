package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"meishi_golang/docs"
	"meishi_golang/pkg/handler"
	"meishi_golang/pkg/repository"
	"meishi_golang/pkg/service"
	"meishi_golang/senti"
	"os"
)

// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	//log to console in json
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//load global config
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %s", err.Error())
	}

	//load .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//swagger config
	docs.SwaggerInfo.Title = "Документация REST API"
	docs.SwaggerInfo.Description = "Описание методов взаимодействия с api."
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = os.Getenv("HOSTSWAGGER")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	//connect to db
	db, err := repository.CuteConnectionDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("filed to init DB: %s", err.Error())
	}

	repos := repository.VeryCuteRepository(db)
	services := service.VeryCuteService(repos)
	handlers := handler.VeryCuteHandler(services)
	svr := new(senti.Server)

	if err := svr.RunMyServerApi(viper.GetString("port"), handlers.RoutingInitialization()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
