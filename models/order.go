// order.go
package models

import "time"

type Order struct {
    OrderID      uint       `gorm:"primary_key" json:"order_id"`
    CustomerName string     `json:"customer_name"`
    OrderedAt    *time.Time `json:"ordered_at"`
    Items        []Item     `json:"items" gorm:"foreignkey:OrderID"`
}

type Item struct {
    ItemID      uint   `gorm:"primary_key" json:"item_id"`
    ItemCode    string `json:"item_code"`
    Description string `json:"description"`
    Quantity    uint   `json:"quantity"`
    OrderID     uint   `json:"order_id"`
}
