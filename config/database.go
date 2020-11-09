package config

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func MongodbConnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin-louis:Test-123@inception.my7v9.mongodb.net/Inception?retryWrites=true&w=majority") // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client

}
