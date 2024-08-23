package main

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-21-latihan-cqrs/config"
	"github.com/ibrahimker/golang-praisindo-advanced/session-21-latihan-cqrs/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func main() {
	r := gin.Default()
	readDB, err := gorm.Open(postgres.Open(config.DBReadDSN), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		slog.Error("error opening read database", slog.String("dsn", config.DBReadDSN), slog.Any("err", err))
		os.Exit(1)
	}
	brokers := []string{config.KafkaBrokerAddress}
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		slog.Error("error when call sarama.NewSyncProducer", slog.Any("error", err))
		os.Exit(1)
	}

	userHandler := handler.NewUserHandler(readDB, producer)

	r.GET("/users", userHandler.GetAllUsers)
	r.POST("/users", userHandler.CreateUser)

	if err = r.Run(":8080"); err != nil {
		slog.Error("run gin error", slog.Any("err", err))
		os.Exit(1)
	}
}
