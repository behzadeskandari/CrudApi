package main

import (
	"CrudApi/controlles"
	"CrudApi/initilizaers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizaers.LoadEnvVariables()

}
func main() {
	initilizaers.ConnectToDB()

	if initilizaers.DB == nil {
		log.Fatal("DB is still nil after ConnectToDB!")
	}
	log.Println("✅ DB initialized successfully")
	fmt.Println("Hello")

	r := gin.Default()
	r.POST("/post", controlles.PostsCreate)
	r.PUT("/post/:id", controlles.PostsUpdate)
	r.GET("/post", controlles.PostsIndex)
	r.GET("/post/:id", controlles.PostsShow)
	r.Run()
}
