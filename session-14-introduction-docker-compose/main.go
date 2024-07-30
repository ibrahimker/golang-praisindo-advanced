package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	dsn := "postgresql://postgres:password@postgres-db:5434/postgres"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("hello world run from docker compose")
	log.Println("hello world run from docker compose 2")
	log.Println(gormDB)
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully ping db")
}
