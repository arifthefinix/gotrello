package Models


type UserWorkspace struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `gorm:"type:int"`
	WorkspaceID uint           `gorm:"type:int"`
	Role        uint		   `gorm:"type:int"`
	CreatedBy   uint 		`gorm:"type:int"`
	UpdatedBy   uint 		`gorm:"type:int"`
}
