package db

import (
	"log"

	"github.com/aaa59891/ticket/config"

	"github.com/aaa59891/ticket/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", config.GetConfig().GetDBConnectString())
	if err != nil {
		log.Fatal(err)
	}
	DB.LogMode(true)
	DB.AutoMigrate(&models.Ticket{})
}
