package models

import (
	"clinic-api/src/database"

	"gorm.io/gorm"
)

var DB = new(database.DBConf).InitDB()

type Polyclinic struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PolyclinicResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
