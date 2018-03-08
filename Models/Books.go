package Models

import (
	"../Config"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllBook(b *[]Book) (err error) {
	//fmt.Println(Config.DB)
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}
