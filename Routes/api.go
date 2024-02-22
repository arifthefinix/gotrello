package Routes

import (
	"github.com/gin-gonic/gin"
	"gotrello/Controllers/WorkspaceController"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	workspaces := router.Group("/workspaces")
	{
		workspaces.GET("/", WorkspaceController.Index)
	}

	return router
}
