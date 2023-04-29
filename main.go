package main

import (
	"ardi_go/auth"
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

	// db.AutoMigrate(&entity.User{})
	// DB = db
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJBdWRpZW5jZSIsImV4cCI6MTY4Mjc4MTE0NSwidXNlcl9pZCI6Mn0.ZqUMMK7u6pCwwQ0mI9WpIb_plhp4QHgVxsDUif7pCmI")
	if err != nil {
		fmt.Println("error")
		fmt.Println("error")
		fmt.Println("error")
	}
	if token.Valid {
		fmt.Println("valid")
		fmt.Println("valid")
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
		fmt.Println("invalid")
	}

	userHandler := handler.NewUseHandler(userService, authService)
	router := gin.Default()

	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()
}
