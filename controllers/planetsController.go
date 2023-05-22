package controllers

import (
	"math/rand"

	"dias375.dev/m/initializers"
	"dias375.dev/m/models"
	"github.com/gin-gonic/gin"
)

func PlanetsCreate(c *gin.Context) {

	var body struct {
		Name    string
		Climate string
		Terrain string
	}

	c.Bind(&body)

	id := rand.Uint64()

	planet := models.Planet{UUID: int64(id), Name: body.Name, Climate: body.Climate, Terrain: body.Terrain}

	result := initializers.DB.Create(&planet)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
