package main

import (
	"crowdfunding/auth"
	"crowdfunding/campaign"
	"crowdfunding/handler"
	"crowdfunding/helper"
	"crowdfunding/payment"
	"crowdfunding/transaction"
	"crowdfunding/user"
	webHanlder "crowdfunding/web/handler"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
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

	userWebHandler := webHanlder.NewUserHandler(userService)
	router := gin.Default()
	router.Use(cors.Default())
	// router.LoadHTMLGlob("web/templates/**/*")
	router.HTMLRender = loadTemplates("./web/templates")
	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	api := router.Group("/api/v1")
	//API ROUTE
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
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	//WEB ROUTE
	router.GET("/users", userWebHandler.Index)
	router.GET("/users/new", userWebHandler.New)
	router.POST("/users", userWebHandler.Create)
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

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
