package database

import (
	"log"

	"golang-test-5-hardware-shop/api/database/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", config.BuildDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
