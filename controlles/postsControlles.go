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
		ctx.JSON(400, gin.H{
			"message": "post not created",
		})
		return
	}
	if result.RowsAffected >= 0 {
		ctx.JSON(400, gin.H{
			"message": "post created",
		})
		return
	}
}
