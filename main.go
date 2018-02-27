package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"net/http"
)

var db *gorm.DB
var err error

type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (b *Book) TableName() string {
	return "book"
}

func main() {
	db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testinger?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Book{})

	// route
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("book", ListBook)
	}

	// running
	r.Run(":1010")
}

func ListBook(c *gin.Context) {
	var book []Book
	if err := db.Find(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		//fmt.Println(book)
		//c.JSON(200, gin.H{"data": book})
		respondJSON(c, 200, book)
	}
}
