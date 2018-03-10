package main

import (
	"./Config"
	"./Models"
	"./Routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/testinger?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
