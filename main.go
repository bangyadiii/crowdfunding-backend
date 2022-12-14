package main

import (
	"backend-crowdfunding/handler"
	"backend-crowdfunding/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error : ", err.Error())
	}
	fmt.Println("berhasil terkoneksi ke database ", db)

	repository := user.NewRepository(db)
	userService := user.NewService(repository)

	userHandler := handler.NewUserHanlder(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/", userHandler.RegisterUser)

	router.Run()	
}
