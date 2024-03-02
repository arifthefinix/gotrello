package Routes

import (
	"github.com/gin-gonic/gin"
	"gotrello/Controllers/WorkspaceController"
	"gotrello/Controllers/UserController"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	workspaces := router.Group("/workspaces")
	{
		workspaces.GET("/", WorkspaceController.Index)
	}

	user := router.Group("/user")
	{
		user.POST("/signup", UserController.Signup)
	}
	return router
}
