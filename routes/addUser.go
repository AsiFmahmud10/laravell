package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"salt/m/db"
	"salt/m/model"
	"salt/m/salti"
)



func RegisterUser(w http.ResponseWriter, r *http.Request){
	//get credentials
	var newUser model.Users 
	json.NewDecoder(r.Body).Decode(&newUser)

	//get salt and the final hashed  value
	salt,hashedSalt  := salti.SaltingHash(salti.GenerateSalt(),newUser.Password)
	

	user := model.Users{Username:newUser.Username,Salt:salt ,Password:hashedSalt}
	//create user
	result := db.GetDb().Create(&user)
	if result.Error != nil {

		log.Fatal("create problem")
	}
	id := fmt.Sprintf("%d",user.ID)
	w.Write([]byte("welcome to add"+id))
}
