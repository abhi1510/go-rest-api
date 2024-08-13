package main

import (
	"github.com/abhi1510/go-rest-api/initializers"
	"github.com/abhi1510/go-rest-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
