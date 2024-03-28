package controllers

import (
	"fmt"
	"net/http"
	"solidithai/orm"

	// "database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func ReadAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
}

func ReadOne(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("User ID:", id) // Log the user ID

	var user orm.User
	if err := orm.Db.Where("id = ?", id).First(&user).Error; err != nil {
		fmt.Println("Error:", err) // Log any errors
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
		return
	}

	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "user not found"})
	}
}

func FindById(c *gin.Context) {
	fmt.Println("FindById")
	var users []orm.User
	orm.Db.Find(&users)
	// fmt.Println("a:", a)
	Username := c.Param("username")
	fmt.Println("Username:", Username)
	// Username := "surawee4"
	// Find user by ID
	fmt.Println("Username:", users)
	for _, user := range users {
		if user.Username == Username {
			c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
