package model

import (
	
	"gorm.io/gorm"
)


func Migrate( Db *gorm.DB){
	
	//auto migrations 
	Db.AutoMigrate(&Users{})

}

