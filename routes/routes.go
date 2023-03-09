package routes

import (
	"github.com/Farheen-cell/UserCRUD/controller"

	"github.com/gin-gonic/gin"
)

// setup	Router ....configer routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("user-api")
	{
		grpl.GET("user", controller.GetUser)
		grpl.POST("user", controller.CreatUser)
		grpl.GET("user/:id", controller.GetUserById)
		grpl.PUT("user/:id", controller.UpdateUser)
		grpl.DELETE("user/:id", controller.DeleteUser)
	}
	return r
}
