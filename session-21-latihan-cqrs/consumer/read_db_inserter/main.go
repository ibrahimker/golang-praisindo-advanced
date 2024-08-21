package main

import (
	"github.com/ibrahimker/golang-praisindo-advanced/session-21-latihan-cqrs/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func main() {
	readDB, err := gorm.Open(postgres.Open(config.DBReadDSN), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		slog.Error("error opening read database", slog.String("dsn", config.DBReadDSN), slog.Any("err", err))
		os.Exit(1)
	}

}
