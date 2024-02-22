package Models

import (
	"gorm.io/gorm"
	"time"
)

type Workspace struct {
	ID          uint           `gorm:"primaryKey"`
	Title	   string         `form:"title" json:"title" binding:"required"`
	Slug       string         `form:"slug" json:"slug" binding:"required"`
	Description string       `form:"description" json:"description" binding:"required"`
	CreatedBy   uint 		`gorm:"type:int"`
	UpdatedBy   uint 		`gorm:"type:int"`
	Sorting    int		   `gorm:"type:int"`
	DeletedBy  uint		   `gorm:"type:int"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
