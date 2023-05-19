package main

import (
	"dias375.dev/m/initializers"
	"dias375.dev/m/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Planet{})
}
