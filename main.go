package main

import (
	account_repositories "clinic-api/src/app/account/repositories"
	admin_repositories "clinic-api/src/app/admin/repositories"
	doctor_repositories "clinic-api/src/app/doctor/repositories"
	polyclinic_repositories "clinic-api/src/app/polyclinic/repositories"
	queue_repositories "clinic-api/src/app/queue/repositories"
	"clinic-api/src/database"
	"clinic-api/src/models"
	"clinic-api/src/routes"
)

func init() {
	new(database.DBConf).InitDB().AutoMigrate(
		models.Patient{},
		account_repositories.User{},
		polyclinic_repositories.Polyclinic{},
		doctor_repositories.Doctor{},
		admin_repositories.Admin{},
		queue_repositories.Queue{},
	)
}

func main() {
	app := routes.New()
	app.Logger.Fatal(app.Start(":8000"))
}
