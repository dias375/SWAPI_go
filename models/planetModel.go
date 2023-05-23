package models

import (
	"gorm.io/gorm"
)

type Planet struct {
	gorm.Model
	ID      string
	Name    string
	Climate string
	Terrain string
}
