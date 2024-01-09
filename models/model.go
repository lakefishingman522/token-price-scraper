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

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
