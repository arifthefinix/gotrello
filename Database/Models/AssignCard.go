package Models

type AssignCard struct {
	ID          uint           `gorm:"primaryKey"`
	CardID      uint           `gorm:"type:int"`
	UserID      uint           `gorm:"type:int"`
}
