package storage

import (
	"github.com/yakuzzaa/GoDone/backendService/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	cfg := config.MustLoad()
	dsn := cfg.DSN()
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

}
