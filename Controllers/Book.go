package Controllers

import (
	"../ApiHelpers"
	"../Models"
	//"fmt"
	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	var book []Models.Book
	/*page := c.DefaultQuery("page", "")
	fmt.Println("page: ", page)*/
	err := Models.GetAllBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func AddNewBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.AddNewBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.PutOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}
