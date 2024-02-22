package Models


type Project struct {
	ID          uint           `gorm:"primaryKey"`
	Title	   string         `gorm:"type:varchar(255)"`
	Slug       string         `gorm:"type:varchar(255);unique"`
	Description string       `gorm:"type:varchar(255)"`
	CreatedBy   uint 		`gorm:"type:int"`
	UpdatedBy   uint 		`gorm:"type:int"`
	Sorting    int		   `gorm:"type:int"`
	WorkspaceID uint 	   `gorm:"type:int"`
}