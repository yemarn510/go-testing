package main

import (
	"test/controllers"

	"test/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/posts", controllers.Create)
	router.GET("/posts", controllers.Get)
	router.GET("/posts/:id", controllers.GetDetail)
	router.PATCH("/posts/:id", controllers.Update)
	router.DELETE("/posts/:id", controllers.Delete)

	router.Run("localhost:8080")
}