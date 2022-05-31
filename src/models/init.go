package models

import "clinic-api/src/database"

var DB = new(database.DBConf).InitDB()
