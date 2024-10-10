package models

import (
	"time"
)

// User struct representing the user table in PostgreSQL
type User struct {
	ID                int64      `gorm:"primaryKey;autoIncrement:false"` // Snowflake ID
	Username          string     `gorm:"size:255;not null;unique"`
	Name              string     `gorm:"size:255"`
	CreatedAt         time.Time  `gorm:"autoCreateTime"`
	UpdatedAt         time.Time  `gorm:"autoUpdateTime"`
	DeletedAt         *time.Time `gorm:"index"` // Soft delete with timestamp
	ProfilePictureURL string     `gorm:"size:255"`
	Email             string     `gorm:"size:255;unique"`
	StatusID          uint       `gorm:"not null"`            // Foreign key to Status table
	Status            Status     `gorm:"foreignKey:StatusID"` // Relationship to Status
}

type Status struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Value string `gorm:"size:50;not null;unique"` // Values like active, inactive, block
}
