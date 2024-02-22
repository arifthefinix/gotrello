package Models



type Comment struct {
	ID 		uint           `gorm:"primaryKey"`
	CardID  uint           `gorm:"type:int"`
	Comment string         `gorm:"type:varchar(255)"`
	CreatedBy uint         `gorm:"type:int"`
}
