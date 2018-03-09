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

func AddNewBook(b *Book) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *Book, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *Book, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	Config.DB.Save(b)
	return nil
}
