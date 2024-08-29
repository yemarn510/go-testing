package main

import (
	"./models"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}


func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.Run("localhost:8080")
}