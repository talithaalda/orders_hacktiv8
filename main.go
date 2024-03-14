package main

import (
	"fmt"
	"log"

	"./handler/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	host := "127.0.0.1"
	port := "5433"
	user := "postgres"
	password := "talitha2712"
	dbname := "order_hactiv8"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	_, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	router := gin.Default()

    // Routes
    router.POST("/orders", handlers.CreateOrder)
    router.GET("/orders/:id", handlers.GetOrder)
    router.PUT("/orders/:id", handlers.UpdateOrder)
    router.DELETE("/orders/:id", handlers.DeleteOrder)

    // Run the server
    log.Fatal(router.Run(":8080"))
}