package models

import (
	"gorm.io/gorm"
)

type Planet struct {
	gorm.Model
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Climate string `json:"Climate"`
	Terrain string `json:"Terrain"`
}
