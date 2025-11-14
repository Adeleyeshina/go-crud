package routes

import (
	"github.com/AdeleyeShina/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	userRoute := r.Group("/api/user")
	{
		userRoute.GET("/", controllers.GetAllUser)
		userRoute.GET("/:id", controllers.GetSingleUser)
		userRoute.POST("/", controllers.CreateUser)
		userRoute.PUT("/:id", controllers.UpdateUser)
		userRoute.DELETE("/:id", controllers.DeleteUser)
	}
}
