package main

import (
	"./Config"
	"./Models"
	"./Routers"
	"fmt"
	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testinger?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	//fmt.Println("ststu")
	Config.DB.AutoMigrate(&Models.Book{})

	// route
	/*	r := gin.Default()
		v1 := r.Group("/v1")
		{
			v1.GET("book", Controllers.ListBook)
		}*/
	r := Routers.SetupRouter()
	// running
	r.Run()
}
