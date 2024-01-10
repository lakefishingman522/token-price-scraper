package db

import (
	_ "github.com/mattn/go-sqlite3"

	"gorm.io/driver/sqlite"

	"github.com/CascadiaFoundation/CascadiaTokenScrapper/models"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

func Init() (*gorm.DB, error) {

	// Connect to database
	db, err := gorm.Open(sqlite.Open("statistics.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	logger := log.NewEntry(log.StandardLogger())
	logger.Info("Connected to Database...")

	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON’T change existing column’s type or delete unused columns to protect your data.
	err = db.AutoMigrate(&models.TokenStatisticsModel{})
	if err != nil {
		panic("failed to create table")
	}

	return db, nil
}
