// handlers.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrder creates a new order
func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, order)
}

// GetOrder retrieves an order by ID
func GetOrder(c *gin.Context) {
    var order models.Order
    id, _ := strconv.Atoi(c.Param("id"))

    if err := db.Preload("Items").First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

// UpdateOrder updates an existing order
func UpdateOrder(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var order models.Order

    if err := db.First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    var updatedOrder models.Order
    if err := c.ShouldBindJSON(&updatedOrder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Model(&order).Updates(updatedOrder)

    c.JSON(http.StatusOK, order)
}

// DeleteOrder deletes an order by ID
func DeleteOrder(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var order models.Order

    if err := db.First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    db.Delete(&order)

    c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
