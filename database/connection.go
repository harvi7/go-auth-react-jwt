package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:root12345@/go-auth-react-jwt"), &gorm.Config{})

	if err != nil {  
		panic("could not connect to the database ")
	}
}