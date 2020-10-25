package main

import (
	"github.com/initiumfund/artis-go/config"
	"github.com/initiumfund/artis-go/models"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDataBase(cfg *config.Config, log *zap.SugaredLogger) (*gorm.DB, error) {
	var db *gorm.DB

	// Developmental database
	if cfg.IsDevelopment() {
		var err error
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}
	// TODO: Production DB

	err := db.AutoMigrate(
		&models.User{},
		&models.UserCostBasis{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
