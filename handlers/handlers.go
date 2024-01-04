package handlers

import (
	"github.com/jinzhu/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}
