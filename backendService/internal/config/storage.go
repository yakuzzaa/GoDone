package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *Config) (*gorm.DB, error) {
	dsn := cfg.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
