package main

import (
	"crowdfunding/auth"
	"crowdfunding/handler"
	"crowdfunding/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepostory := user.NewRepository(db)
	userService := user.NewService(userRepostory)
	authService := auth.NewService()
	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.BZcBBLhOhjA9ojwmRNMLx7x0IR83QyTeiH48psbhKLI")
	if err != nil {
		fmt.Println("error")
	}
	if token.Valid {
		fmt.Println("valid")
	}
	userhandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userhandler.RegisterUser)
	api.POST("/sessions", userhandler.Login)
	api.POST("/email_checkers", userhandler.CheckEmailAvailability)
	api.POST("/avatars", userhandler.UploadAvatar)
	router.Run()

}
