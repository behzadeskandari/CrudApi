package main

import (
	"CrudApi/initilizaers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizaers.LoadEnvVariables()
	initilizaers.ConnectToDB()
}
func main() {
	fmt.Println("Hello")

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
