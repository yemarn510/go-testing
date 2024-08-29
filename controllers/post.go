package controllers

import (
	"net/http"

	"test/models"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostInput struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func Create(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); 
	
	err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
 	}

	post := models.Post{ Title: input.Title, Content: input.Content } 
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{ "data": post })

}

func Get(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts })
}

func GetDetail(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; 
	
	err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func Update(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not fo"})
		return
	}

	var input UpdatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatePost := models.Post{ Title: input.Title, Content: input.Content }

	models.DB.Model(&post).Updates(&updatePost)
	c.JSON(http.StatusOK, gin.H{"data": post })
}

func Delete(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ? ", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": "Post deleted successfully"})
}