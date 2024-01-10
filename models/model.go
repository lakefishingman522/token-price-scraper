package models

import "time"

type TokenStatisticsModel struct {
	ID        uint64 `gorm:"primaryKey"`
	TimeStamp uint64 `gorm:"not null"`

	P360 float64 `gorm:"not null"`
	P180 float64 `gorm:"not null"`
	P90  float64 `gorm:"not null"`
	P30  float64 `gorm:"not null"`
	P14  float64 `gorm:"not null"`
	P7   float64 `gorm:"not null"`
	P1   float64 `gorm:"not null"`

	PreviousAveragePrice360 float64 `gorm:"not null"`
	PreviousAveragePrice180 float64 `gorm:"not null"`
	PreviousAveragePrice90  float64 `gorm:"not null"`
	PreviousAveragePrice30  float64 `gorm:"not null"`
	PreviousAveragePrice14  float64 `gorm:"not null"`
	PreviousAveragePrice7   float64 `gorm:"not null"`
	PreviousAveragePrice1   float64 `gorm:"not null"`

	CurrentAveragePrice360 float64 `gorm:"not null"`
	CurrentAveragePrice180 float64 `gorm:"not null"`
	CurrentAveragePrice90  float64 `gorm:"not null"`
	CurrentAveragePrice30  float64 `gorm:"not null"`
	CurrentAveragePrice14  float64 `gorm:"not null"`
	CurrentAveragePrice7   float64 `gorm:"not null"`
	CurrentAveragePrice1   float64 `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
