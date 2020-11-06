package config

import (
	"fmt"
	"log"
	"os"

	"myapp/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

//ConnectDB connect GO to database
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.InitLog(),
	})

	if err != nil {
		log.Print(err)
		panic("Failed to connect database")
	}

	return db
}
