package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var database_setting_main = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", database_setting_main)

	if err != nil {
		log.Print(err)
		panic("Failed to connect database")
	}

	db.LogMode(true)
	db.SingularTable(true)

	return db
}
