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
