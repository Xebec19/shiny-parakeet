package db

import (
	"log"
	"time"

	"github.com/Xebec19/shiny-parakeet/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("can not connect to database ", err)
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal("can not create a database pool ", err)
	}
	log.Printf("Database connected successfully")
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
