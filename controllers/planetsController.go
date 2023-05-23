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

	result := initializers.DB.FirstOrCreate(&planet)

	//result := initializers.DB.Create(&planet)

	c.JSON(200, result)
}
