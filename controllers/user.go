package controllers

import (
	"strconv"

	"github.com/THEMEGANAME/workshop-golang/middlewares"

	"github.com/THEMEGANAME/workshop-golang/models"
	"github.com/gin-gonic/gin"
)

// ResgisterUserEndPoints - this is end point of users
func ResgisterUserEndPoints(userRouter *gin.RouterGroup) {

	userRouter.GET("/", getUsers)

	userRouter.Use(middlewares.Authentication())
	{
		userRouter.GET("/:id", getUser)
		userRouter.POST("/create", createUser)
		userRouter.POST("/update", updateUser)
		userRouter.POST("/delete", deleteUser)
	}

}

// getUsers - this is function get users
func getUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "Error get users !",
		})
	}
	c.JSON(200, gin.H{
		"msg":   "get users success",
		"users": users,
	})
}

// getUser - this is function get user detail
func getUser(c *gin.Context) {
	var user models.User
	uid, _ := strconv.Atoi(c.Param("id"))
	user.ID = uint(uid)
	err := user.GetUserDetail()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "get user unsuccess",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "get user success",
		"data": user,
	})
}

// createUser - this is function create  user
func createUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	err = user.CreateUser()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "create user unsuccess",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "create user success",
	})
}

// updateUser - this is function update  user
func updateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	err = user.UpdateUser()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "update user unsuccess",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "update user success",
	})
}

// deleteUser - this is function delete  user
func deleteUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bad request",
		})
		return
	}
	err = user.DeleteUser()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "delete user unsuccess",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "delete user success",
	})
}
