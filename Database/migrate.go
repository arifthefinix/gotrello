package Database

import "gotrello/Database/Models"

func init() {
	ConnectToDatabase()
}

func Migrate() error{
	DB.AutoMigrate(&Models.AssignCard{})
	DB.AutoMigrate(&Models.Board{})
	DB.AutoMigrate(&Models.BoardUser{})
	DB.AutoMigrate(&Models.Card{})
	DB.AutoMigrate(&Models.Comment{})
	DB.AutoMigrate(&Models.Project{})
	DB.AutoMigrate(&Models.ProjectUser{})
	DB.AutoMigrate(&Models.User{})
	DB.AutoMigrate(&Models.UserWorkspace{})
	DB.AutoMigrate(&Models.Workspace{})

	return nil
}