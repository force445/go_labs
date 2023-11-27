package database

import (
	"fmt"
	"go-posgres/config"
	"go-posgres/model"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port) //What is Sprintf? https://golang.org/pkg/fmt/#Sprintf
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error connect database", err)
		os.Exit(1)
	}
	log.Println("Database connected")
	db.Logger = db.Logger.LogMode(logger.Info)
	log.Println("run migration")
	db.AutoMigrate(&model.Member{})
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.CreditCard{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.Author{})
	db.AutoMigrate(&model.Publisher{})

	DB = Dbinstance{
		Db: db,
	}
}
