package main

import (
	"dias375.dev/m/controllers"
	"dias375.dev/m/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/planets/", controllers.PlanetsCreate)
	r.Run()
}
