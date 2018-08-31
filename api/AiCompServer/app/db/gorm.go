package db

import (
	"AiCompServer/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() {
	dbtype := revel.Config.StringDefault("db.type", "sqlite3")
	dbargs := revel.Config.StringDefault("db.args", "./db.sqlite3")
	for {
		db, err := gorm.Open(dbtype, dbargs)
		if err != nil {
			log.Println("Failed to connect to database: %v\n", err)
		} else {
			db.DB()
			db.AutoMigrate(&models.User{})
			db.AutoMigrate(&models.Challenge{})
			db.AutoMigrate(&models.Answer{})
			db.AutoMigrate(&models.PreAnswer{})
			db.AutoMigrate(&models.Notification{})
			db.AutoMigrate(&models.Setting{})
			DB = db
			break
		}
		time.Sleep(1 * time.Second)
	}
}
