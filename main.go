package main

import (
	"fmt"
	"log"
	"tidy/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db *gorm.DB
func main() {
	host := "localhost"
	port := "5433"
	user := "postgres"
	password := "talitha2712"
	dbname := "orders_hactiv8"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	handler.SetDB(db)
	router := gin.Default()

    // Routes
	router.GET("/orders", handler.GetAllOrders)
    router.POST("/orders", handler.CreateOrder)
    router.GET("/orders/:id", handler.GetOrder)
    router.PUT("/orders/:id", handler.UpdateOrder)
    router.DELETE("/orders/:id", handler.DeleteOrder)

    // Run the server
    log.Fatal(router.Run(":8080"))
}