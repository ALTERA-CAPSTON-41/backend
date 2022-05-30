package main

import (
	"clinic-api/src/database"
	"clinic-api/src/routes"
)

func init() {
	new(database.DBConf).InitDB().Migrate()
}

func main() {
	app := routes.New()
	app.Logger.Fatal(app.Start(":8000"))
}
