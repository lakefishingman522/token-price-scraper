package db

import (
	"fmt"
	"os"

	_ "github.com/CascadiaFoundation/CascadiaTokenScrapper/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() (*gorm.DB, error) {
	// Read hostname, password, dbname and username from environment variables
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, dbName, password)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	// db.AutoMigrate(&models.Address{}, &models.Query{})
	// db.Model(&models.Address{}).AddForeignKey("query_id", "queries(query_id)", "CASCADE", "CASCADE")
	return db, nil
}
