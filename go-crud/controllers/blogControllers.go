package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/misbah-bp/go-crud/initializers"
	"github.com/misbah-bp/go-crud/models"
)

func BlogCreate(c *gin.Context) {

	// get data of request body

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// create a blog
	blog := models.Blog{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&blog)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return responce

	c.JSON(200, gin.H{
		"Blog":   blog,
		"sucess": "Blog created successfully!",
	})
}

func BlogGet(c *gin.Context) {

	// Get the all blogs
	var blogs []models.Blog
	initializers.DB.Find(&blogs)

	// responce blog
	c.JSON(200, gin.H{
		"Blog": blogs,
	})
}

func SingelBlog(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// Get the all blogs
	var blog models.Blog
	result := initializers.DB.First(&blog, id)

	// Check if the blog was found
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "No blog found"})
		return
	}
	// responce blog
	c.JSON(200, gin.H{
		"Blog": blog,
	})
}

func BlogUpdate(c *gin.Context) {
	// Get the if from url
	id := c.Param("id")

	// Gte the data Off request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// find the Post where updating
	var blog models.Blog
	initializers.DB.First(&blog, id)

	// Update it
	initializers.DB.Model(&blog).Updates(models.Blog{
		Title: body.Title,
		Body:  body.Body,
	})

	// Responce with it
	c.JSON(200, gin.H{
		"Blog": blog,
	})
}

func DeleteBlog(c *gin.Context) {
	// Get the if from url
	id := c.Param("id")

	var blog models.Blog
	result := initializers.DB.Delete(&blog, id)

	// Check if the blog was found
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Not delete"})
		return
	}

	// Responce with it
	c.JSON(200, gin.H{

		"sucess": "Blog Deleted",
	})

}
