package main

import (
	"fmt"
	"karyaoneplus/auth"
	"karyaoneplus/handler"
	"karyaoneplus/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/karyaoneplus?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.o_gE9qVP0kL6inwLg7ndxR_OpeDloDX727aLX_dU9Jg")
	if err != nil {
		fmt.Println("Error")
		fmt.Println("Error")
		fmt.Println("Error")
	}

	if token.Valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatars)
	router.Run()
}
