package main

import (
	"fmt"
	"solidithai/orm"
	routes "solidithai/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

type Data struct {
	value string
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()
	// initail route
	router := gin.Default()

	// controllers.ReccomendPost()
	// use middleware
	router.Use(CORSMiddleware())
	// router.Use(orm.Verify)

	routes.UserRoute(router)
	// routes.QuestRoute(router)
	// routes.PostRoutes(router)
	// routes.CommentRoutes(router)
	// routes.LikeRoutes(router)
	// routes.BookMarkRoutes(router)
	// routes.TestRoute(router)
	// // routes.ReportRoutes(router)
	// routes.MockingRoute(router)

	// configue on port 3000
	router.Run("0.0.0.0:8080")

	// r := gin.Default()
	// r.Use(cors.Default())
	// r.POST("/register", AuthController.Register)
	// r.POST("/login", AuthController.Login)
	// authorized := r.Group("/users", middleware.JWTAuthen())
	// authorized.GET("/readall", UserController.ReadAll)
	// r.Run("localhost:3000")

}
