package Models

type ProjectUser struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `gorm:"type:int"`
	ProjectID   uint           `gorm:"type:int"`
	Role        uint		   `gorm:"type:int"`
	CreatedBy   uint 		`gorm:"type:int"`
	UpdatedBy   uint 		`gorm:"type:int"`
}

