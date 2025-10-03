package controlles

import (
	"CrudApi/initilizaers"
	model "CrudApi/models"
	"log"

	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	log.Println("body is %T\n", body)
	ctx.Bind(&body)

	log.Println("body222 is %T\n", body)

	post := model.Post{Title: body.Title, Body: body.Body}

	result := initilizaers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}
}
func PostsIndex(ctx *gin.Context) {

	var post []model.Post

	initilizaers.DB.Find(&post)

	ctx.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post []model.Post

	initilizaers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsUpdate(ctx *gin.Context) {
	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Fetch existing post by ID (example; adjust as needed)
	id := ctx.Param("id") // Assuming route: PUT /post/:id
	var existingPost model.Post
	if err := initilizaers.DB.First(&existingPost, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	// Update: Use Updates with pointer to temp struct (or &existingPost directly)
	updateData := model.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initilizaers.DB.Model(&existingPost).Updates(&updateData) // Key: Updates(&struct)

	if result.Error != nil {
		log.Printf("❌ Update failed: %v", result.Error)
		ctx.JSON(400, gin.H{"error": "Update failed: " + result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(400, gin.H{"error": "No changes applied"})
		return
	}

	log.Printf("✅ Updated post ID: %d", existingPost.ID)
	ctx.JSON(200, gin.H{
		"message": "Post updated",
		"post":    existingPost, // Returns updated with timestamps
	})

}
