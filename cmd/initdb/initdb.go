package main

import (
	"fmt"
	"github.com/charlesren/eagle/pkg/config"
	"github.com/charlesren/eagle/pkg/order"
	"github.com/charlesren/eagle/pkg/portfolio"
)

func main() {
	// connect to database
	db := config.GetDB()
	// migrate schema
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&portfolio.OpenInterest{})
	fmt.Println("migrate schema successfully!!!")
}