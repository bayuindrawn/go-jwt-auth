package structs

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}

type Users struct {
	gorm.Model
	Username    string `faker:"username"`
	Password    string `faker:"password"`
	Email       string `faker:"email"`
	Token       string `faker:"jwt"`
	PhoneNumber string `faker:"phone_number"`
}
