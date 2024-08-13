package controllers

import (
	"github.com/abhi1510/go-rest-api/initializers"
	"github.com/abhi1510/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	// Retrieve all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Send Response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func CreatePost(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{
		Title: body.Title, 
		Body: body.Body}
	result := initializers.DB.Create(&post)

	// Send Response
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetSinglePost(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Retrieve the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Send Response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Retrieve the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Get data off req body
	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)

	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	// Send Response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Retrieve the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Delete the post
	initializers.DB.Model(&post).Delete(&models.Post{}, id)

	// Send Response
	c.Status(200)
}
