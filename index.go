package main

import (
	"net/http"
	"log"
	"os"
    "github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

func HomeHandler( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("welcome"))
}

func main(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	
	PORT := os.Getenv("PORT")


	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe("0.0.0.0:"+PORT,r)
}