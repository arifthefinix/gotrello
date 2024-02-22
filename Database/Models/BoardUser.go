package Models

type BoardUser struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `gorm:"type:int"`
	BoardID     uint           `gorm:"type:int"`
	Role        uint		   `gorm:"type:int"`
	CreatedBy   uint 		`gorm:"type:int"`
	UpdatedBy   uint 		`gorm:"type:int"`
}
