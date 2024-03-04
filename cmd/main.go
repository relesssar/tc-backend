package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os/signal"
	"syscall"
	tc "tc_kaztranscom_backend_go"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"tc_kaztranscom_backend_go/pkg/handler"
	"tc_kaztranscom_backend_go/pkg/repository"
	"tc_kaztranscom_backend_go/pkg/service"
)

// @title Total Control
// @version 1.0
// @description API Сервер для Программы Total Control

// @contact.name Серышев А.В., Яценко К.Ю.
// @contact.email 7773785631@mail.ru,K.Yatsenko@kaztranscom.kz

// @host tc.kaztranscom.kz:9000
// @BasePath /
// @Schemes https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//TODO тут Хард кор пути файла настроек на сервере
	env_path:="/home/ubuntu/go/src/tc_kaztranscom_backend_go/.env"

	//TODO тут локальный файл
	//env_path := "/home/user/work/src/tc_kaztranscom_backend_go/.env"

	if err := godotenv.Load(env_path); err != nil {
		logrus.Fatalf("Ошибка чтения конфиг файла: %s %s", env_path, err.Error())
	}

	SetMode(os.Getenv("GIN_MODE"))

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	dbContract, err := repository.NewMysqlDBContracts(repository.ConfigContracts{
		Host:     os.Getenv("DB_HOST_CONTRACTS"),
		Port:     os.Getenv("DB_PORT_CONTRACTS"),
		Username: os.Getenv("DB_USER_CONTRACTS"),
		DBName:   os.Getenv("DB_NAME_CONTRACTS"),
		Password: os.Getenv("DB_PASS_CONTRACTS"),
	})
	if err != nil {
		//logrus.Fatalf("failed to initialize db: %s", err.Error())
		logrus.Infof("failed to initialize db kcrm: %s", err.Error())
	}

	repos := repository.NewRepository(db, dbContract)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(tc.Server)
	go func() {
		//if err := srv.Run(os.Getenv("API_PORT"), handlers.InitRoutes()); err != nil {
		if err := srv.Run(os.Getenv("API_PORT"), handlers.InitRoutes()); err != nil {
			logrus.Error(err.Error())
		}
	}()

	logrus.Print("Total Control Rest Full Api, запущен на порту " + os.Getenv("API_PORT"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Total Control Rest Full Api, Выключен")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func SetMode(mode string) {
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		panic("mode unavailable. (debug, release, test)")
	}
}
