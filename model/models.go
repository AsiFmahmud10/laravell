package model

import "gorm.io/gorm"



type Users struct{
	gorm.Model
	Username  string	`json : "username"`
	Salt	  string	`json :	""`
	Password  string	`json : "Password"`
}