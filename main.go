package main

import (
	// "fmt"
	"food_delivery/common"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	common.SQLModel
	Name    string `gorm:"column:name;"`
	Address string `gorm:"column:addr;"`
}

func main() {
	os.Setenv("DB_CONN_STR", "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for detail7
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Println(db, err)
}
