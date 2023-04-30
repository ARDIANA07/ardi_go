package main

import (
	"ardi_go/auth"
	"ardi_go/campaign"
	"ardi_go/handler"
	"ardi_go/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/ardi_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection to Database is good")

	// db.AutoMigrate(&campaign.CampaignImage{})
	// DB = db
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	campaigns, err := campaignRepository.FindByUserId(1)

	fmt.Println("debug")
	fmt.Println(len(campaigns))
	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
		fmt.Println("jumlah gambar")
		fmt.Println(len(campaign.CampaignImages))
		fmt.Println(campaign.CampaignImages[0].FileName)

	}

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUseHandler(userService, authService)
	router := gin.Default()

	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()
}

// func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if !strings.Contains(authHeader, "Bearer") {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		//bearer tokentokentoke
// 		tokenSring := ""
// 		arrayToken := strings.Split(authHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenSring = arrayToken[1]
// 		}

// 		token, err := authService.ValidateToken(tokenSring)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		claim, ok := token.Claims.(jwt.MapClaims)

// 		if !ok || !token.Valid {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		userID := int(claim["user_id"].(float64))

// 		user, err := userService.GetUserByID(userID)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		c.Set("CurrentUser", user)

// 	}
// }
