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

func HomeHandler( w http.ResponseWriter, r *http.Request){
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
	


	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe("0.0.0.0:"+PORT,r)
}