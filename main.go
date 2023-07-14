package main

import (
	"crowdfunding/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepostory := user.NewRepository(db)
	userService := user.NewService(userRepostory)
	userInput := user.RegisterUserInput{}
	userInput.Name = "test simpan dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "test occupation"
	userInput.Password = "password"
	userService.RegisterUser(userInput)

}
