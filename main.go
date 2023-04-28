package main

import (
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

	userHandler := handler.NewUseHandler(userService)
	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
}