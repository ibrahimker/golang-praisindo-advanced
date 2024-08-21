package main

import (
	"github.com/ibrahimker/golang-praisindo-advanced/session-21-latihan-cqrs/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func main() {
	writeDB, err := gorm.Open(postgres.Open(config.DBWriteDSN), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		slog.Error("error opening write database", slog.String("dsn", config.DBWriteDSN), slog.Any("err", err))
		os.Exit(1)
	}
}
