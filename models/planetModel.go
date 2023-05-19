package models

import (
	"gorm.io/gorm"
)

type Planet struct {
	gorm.Model
	Uuid    int64
	Name    string
	Climate string
	Terrain string
}
