package main

import (
	"crowdfunding/handler"
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
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

	userhandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userhandler.RegisterUser)
	api.POST("/sessions", userhandler.Login)
	router.Run()

}
