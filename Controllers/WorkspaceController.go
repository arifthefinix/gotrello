package Controllers/WorkspaceController

import (
	"github.com/gin-gonic/gin"
	"gotrello/Database"
	"gotrello/Database/Models"
	"net/http"
)

func Index(c *gin.Context) {
	var workspaces []Models.Workspace

	if err :=Database.DB.Find(&workspaces).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": workspaces})
}

func Create(c *gin.Context) {
	var workspace Models.Workspace

	if err := c.ShouldBindJSON(&workspace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := Database.DB.Create(&workspace).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not created!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sucess": "Workspace created successfully!"})
}