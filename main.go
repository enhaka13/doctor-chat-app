package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "freedb_march:E6gRgU$Ca#aKCBn@tcp(sql.freedb.tech:3306)/freedb_doctorchat?parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("database connected...")
	}
}
