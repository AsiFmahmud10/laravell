package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"salt/m/db"
	"salt/m/model"
	"salt/m/salti"
)

func Login(w http.ResponseWriter, r *http.Request){
	
	db := db.GetDb()
	var user model.Users
	var requestedUser model.Users
	
	json.NewDecoder(r.Body).Decode(&requestedUser)
	//query in db
	result := db.Where("Username = ?",requestedUser.Username).First(&user)
	

	if result.Error != nil{
		w.Write([]byte("register plz"));
	}

	
	_ ,hashedStrsalt :=  salti.SaltingHash(user.Salt, requestedUser.Password)

    if user.Password == hashedStrsalt{
		w.Write([]byte("successfully loged in "))

	}
	fmt.Printf("userdb %s\n", user.Password)
	fmt.Printf("request hash %s\n", hashedStrsalt)
}