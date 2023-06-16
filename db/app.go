package db

import (
	"fmt"
	"log"
	"os"
	"salt/m/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func  GetDb() *gorm.DB {
	return Db
}


func Connect(){

	// load the env variable
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
    
	// Connection to db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",DB_HOST,DB_USER,DB_PASSWORD ,DB_NAME,DB_PORT)

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		 log.Fatal(err)
		 print(err)	
	 }
 
	 //model migration
	 model.Migrate(Db)



	return
}

