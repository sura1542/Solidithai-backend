package routers

import (
	controllers "solidithai/controllers"
	middleware "solidithai/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {
	v1 := route.Group("v1/users")
	authorized := route.Group("v1/auth", middleware.JWTAuthen())
	authorized.GET("/readall", controllers.ReadAll)
	v1.GET("/findbyid/:username", controllers.FindById)
	v1.POST("/register", controllers.Register)
	v1.POST("/login", controllers.Login)
	// v1.POST("/signIn/google", controllers.SignInGoogle)
	// v1.GET("/id", controllers.GetUserById)
	// v1.PUT("/pet/id", controllers.NamingPet)
	// v1.DELETE("/id", controllers.DeleteUser)

}
