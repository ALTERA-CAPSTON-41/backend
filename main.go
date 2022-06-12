package main

import (
	"clinic-api/src/database"
	"clinic-api/src/models"
	"clinic-api/src/routes"
)

func init() {
	new(database.DBConf).InitDB().AutoMigrate(
		models.Patient{},
		models.User{},
		models.Polyclinic{},
		models.Doctor{},
		models.Queue{},
	)
}

func main() {
	app := routes.New()
	app.Logger.Fatal(app.Start(":8000"))
}
