package database

import (
	"clinic-api/src/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConf struct{ *gorm.DB }

func (DB *DBConf) InitDB() *DBConf {
	config, _ := configs.LoadServerConfig(".")
	dsn := config.ConnectionString

	conn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &DBConf{conn}
}

func (DB *DBConf) Migrate() {
	DB.AutoMigrate()
}

func (DB *DBConf) Demigrate() {
	DB.Migrator().DropTable()
}
