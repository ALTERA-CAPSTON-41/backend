package main

import (
	admin_repositories "clinic-api/src/app/admin/repositories"
	polyclinic_repositories "clinic-api/src/app/polyclinic/repositories"
	"clinic-api/src/database"
	"clinic-api/src/models"
	"clinic-api/src/routes"
)

func init() {
	new(database.DBConf).InitDB().AutoMigrate(
		models.Patient{},
		models.User{},
		polyclinic_repositories.Polyclinic{},
		models.Doctor{},
		admin_repositories.Admin{},
		models.Queue{},
	)
}

func main() {
	app := routes.New()
	app.Logger.Fatal(app.Start(":8000"))
}
