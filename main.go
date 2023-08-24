package main

import (
	"crowdfunding/auth"
	"crowdfunding/campaign"
	"crowdfunding/handler"
	"crowdfunding/helper"
	"crowdfunding/payment"
	"crowdfunding/transaction"
	"crowdfunding/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	//repository
	userRepostory := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	//service
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)
	userService := user.NewService(userRepostory)
	paymenService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymenService)
	//TESTING

	//handler
	userhandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")
	//ROUTE Users
	api.POST("/users", userhandler.RegisterUser)
	api.POST("/sessions", userhandler.Login)
	api.POST("/email_checkers", userhandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleWare(authService, userService), userhandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleWare(authService, userService), userhandler.FetchUser)

	//Route campaign
	api.POST("/campaigns", authMiddleWare(authService, userService), campaignHandler.CreateCampaign)
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.PUT("/campaigns/:id", authMiddleWare(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleWare(authService, userService), campaignHandler.UploadImage)

	//Route Transaction Campaign
	api.GET("/campaigns/:id/transactions", authMiddleWare(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddleWare(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleWare(authService, userService), transactionHandler.CreateTransaction)

	router.Run()

}
func authMiddleWare(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unautorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unautorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unautorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userID)
		if err != nil {
			response := helper.APIResponse("Unautorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
