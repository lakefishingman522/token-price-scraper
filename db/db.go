package db

import (
	_ "github.com/mattn/go-sqlite3"

	"gorm.io/driver/sqlite"

	_ "github.com/CascadiaFoundation/CascadiaTokenScrapper/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	return db, nil
}
