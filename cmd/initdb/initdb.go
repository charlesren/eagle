package main

import (
	"fmt"
	"github.com/charlesren/eagle/pkg/config"
	"github.com/charlesren/eagle/pkg/order"
)

func main() {
	// connect to database
	db := config.GetDB()
	// migrate schema
	db.AutoMigrate(&order.Order{})
	fmt.Println("migrate schema successfully!!!")
}