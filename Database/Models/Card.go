package Models

import (
	"time"
)

type Card struct {
	ID          uint           `gorm:"primaryKey"`
	Title	   string         `gorm:"type:varchar(255)"`
	Slug       string         `gorm:"type:varchar(255);unique"`
	Description string       `gorm:"type:varchar(255)"`
	WorkspaceID uint 	   `gorm:"type:int"`
	ProjectID   uint 	   `gorm:"type:int"`
	BoardID     uint 	   `gorm:"type:int"`
	Sorting    int		   `gorm:"type:int"`
	Image string         `gorm:"type:varchar(255)"`
	Deadline time.Time
	Duration time.Time
	CreatedBy   uint 		`gorm:"type:int"`
	UpdatedBy   uint 		`gorm:"type:int"`
	DeleredBy  uint		   `gorm:"type:int"`
}
