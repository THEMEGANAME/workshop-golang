package main

import (
	"github.com/THEMEGANAME/workshop-golang/controllers"
	"github.com/THEMEGANAME/workshop-golang/models"
	"github.com/gin-gonic/gin"
)

func init() {
	models.InitDB()
}

func main() {
	router := gin.Default()

	controllers.ResgisterUserEndPoints(router.Group("/user"))

	router.Run()
}
