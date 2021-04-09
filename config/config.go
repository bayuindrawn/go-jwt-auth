package config

import (
	"github.com/bayuindrawn/go-auth-jwt/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BDInit() (*gorm.DB, error) {
	dsn := "root" + ":" + "" + "@(" + "localhost" + ")/" + "godb" + "?charset=" + "utf8" + "&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(structs.Person{})
	db.AutoMigrate(structs.Users{})

	return db, nil
}
