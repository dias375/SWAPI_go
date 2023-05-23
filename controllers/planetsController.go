package controllers

import (
	"dias375.dev/m/initializers"
	"dias375.dev/m/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func PlanetsCreate(c *gin.Context) {

	var body struct {
		Name    string
		Climate string
		Terrain string
	}

	c.Bind(&body)

	id := uuid.NewV4().String()

	planet := models.Planet{ID: id, Name: body.Name, Climate: body.Climate, Terrain: body.Terrain}

	initializers.DB.Save(planet)

	c.JSON(200, gin.H{"Planet": planet})
}
