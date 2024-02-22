package Models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID 	  uint           `gorm:"primaryKey"`
	Name  string         `gorm:"type:varchar(255)"`
	Email string         `gorm:"type:varchar(255);unique"`
	Phone string         `gorm:"type:varchar(20);unique"`
	Password string      `gorm:"type:varchar(255)"`
	Photo string         `gorm:"type:varchar(255)"`
	Username string      `gorm:"type:varchar(255);unique"`
	AccessTokens string  `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}