package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Username  string	`json : "username"`
	Password  string	`json : "Password"`
}

func HomeHandler( w http.ResponseWriter, r *http.Request){
	
	w.Write([]byte("welcome"))
}


var Db *gorm.DB

func  GetDb() *gorm.DB {
	return Db
}

func AddUser(w http.ResponseWriter, r *http.Request){
	user := User{Username:"kasim",Password:"123123"}
	Db.Create(&user)
	id := fmt.Sprintf("%d",user.ID)
	w.Write([]byte("welcome"+ id))
}


func main(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	
	PORT := os.Getenv("PORT")
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",DB_HOST,DB_USER,DB_PASSWORD ,DB_NAME,DB_PORT)
	
    print(dsn);

	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    //Db = db;
	if err != nil{
		print("Nowww")	
	}
	print("hello")


	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/add", AddUser)
	http.ListenAndServe("0.0.0.0:"+PORT,r)
}