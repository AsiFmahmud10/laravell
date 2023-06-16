package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"salt/m/db"
	"salt/m/routes"
	
)


func main(){

	db.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/add", routes.RegisterUser)
	r.HandleFunc("/login", routes.Login)
	log.Fatal(http.ListenAndServe("0.0.0.0:80",r))
	

}




func HomeHandler( w http.ResponseWriter, r *http.Request){
	
	w.Write([]byte("welcome"))
}