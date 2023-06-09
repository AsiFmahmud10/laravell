package main

import (
	"net/http"
	"log"
	"os"
    "github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Username  string	`json : "username"`
	Password  string	`json : "Password"`
}

func HomeHandler( w http.ResponseWriter, r *http.Request){
	Db.Create(&User{Username:"kasim",Password:"123123"})
	var  new User
	Db.First(&new)
	w.Write([]byte("welcome"))
}
var Db *gorm.DB

func  GetDb() *gorm.DB {
	return Db
}

func main(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	
	PORT := os.Getenv("PORT")
	dsn := "root:udeXsP6p2Xokp1dVkfNT@tcp(containers-us-west-189.railway.app:7476)/railway?charset=utf8mb4&parseTime=True&loc=Local"
  	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    Db = db;
	if err != nil{
		print("Nowww")	
	}
	print("hello")


	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe("0.0.0.0:"+PORT,r)
}