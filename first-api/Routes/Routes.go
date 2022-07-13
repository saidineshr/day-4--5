package Routes

import (
	"first-api/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user-api/user", Controllers.GetUsers)
	r.POST("/user-api/user", Controllers.CreateUser)
	r.GET("/user-api/user/:id", Controllers.GetUserByID)
	r.PUT("/user-api/user/:id", Controllers.UpdateUser)
	r.DELETE("/user-api/user/:id", Controllers.DeleteUser)

	return r
}
