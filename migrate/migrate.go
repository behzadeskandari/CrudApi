package main

import (
	"CrudApi/initilizaers"
	model "CrudApi/models"
)

func init() {
	initilizaers.LoadEnvVariables()
	initilizaers.ConnectToDB()
}

func main() {
	initilizaers.DB.AutoMigrate(&model.Post{})
}
