package main

import (
	"github.com/gin-gonic/gin"
	"github.com/THEMEGANAME/workshop-golang/models"
	"strconv"
)

func init() {
	models.InitDB()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Ping pong !",
		})
	})

	router.GET("/user",func(c *gin.Context){
		users, err := models.GetUsers()
		if err!=nil {
			c.JSON(500,gin.H{
				"msg":"Error get users !",
			})
		}
		c.JSON(200,gin.H{
			"msg":"get users success",
			"users":users,
		})
	})

	router.POST("/user/create",func(c *gin.Context){
		var user models.User
		err := c.BindJSON(&user)
		if err !=nil {
			c.JSON(400,gin.H{
				"msg":"bad request",
			})
			return
		}
		err = user.CreateUser()
		if err !=nil {
			c.JSON(500,gin.H{
				"msg":"create user unsuccess",
			})
			return
		}

		c.JSON(200,gin.H{
			"msg":"create user success",
		})

	})


	router.GET("/user/:id",func(c *gin.Context){
		var user models.User
		uid ,_ := strconv.Atoi(c.Param("id"))
		user.ID = uint(uid)
		err := user.GetUserDetail()
		if err != nil{
			c.JSON(500,gin.H{
				"msg":"get user unsuccess",
			})
			return
		}
		c.JSON(200,gin.H{
			"msg":"get user success",
			"data":user,
		})
	})

	router.POST("/user/update",func(c *gin.Context){
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400,gin.H{
				"msg":"bad request",
			})
			return
		}
		err = user.UpdateUser()
		if err != nil{
			c.JSON(500,gin.H{
				"msg":"update user unsuccess",
			})
			return
		}
		c.JSON(200,gin.H{
			"msg":"update user success",
		})
	})

	router.POST("/user/delete",func(c *gin.Context){
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400,gin.H{
				"msg":"bad request",
			})
			return
		}
		err = user.DeleteUser()
		if err != nil{
			c.JSON(500,gin.H{
				"msg":"delete user unsuccess",
			})
			return
		}
		c.JSON(200,gin.H{
			"msg":"delete user success",
		})
	})





	router.Run()
}
