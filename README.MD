# Api go simple
It is a just simple RESTful API with Go using:
1. **Gin Framework**
2. **Gorm** 

## Installation & Run
```bash
# Download this project
$ go get github.com/dedidot/gorm-gin

# Download Gin Framework
$ go get github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm
```

Setting DB in main.go
```go
Config.DB, err = gorm.Open("mysql", "DBUSERNAME:DBPASSWORD@tcp(127.0.0.1:3306)/DBNAME?charset=utf8&parseTime=True&loc=Local")
```

## Structure
```
├── Models
│   ├── Book.go // Book models
|	├── Scheme.go // Book struct and tabel
├── Config
│   └── Database.go // Global DB
├── Controllers
│   └── Book.go // Book Controller
├── ApiHelpers
│   └── Response.go // response function
├── Routers
|   └── Routers.go // Routers
└── main.go
```

## API

#### /book
* `GET` : Get all book
* `POST` : Create a new book

#### /book/:id
* `GET` : Get a book
* `PUT` : Update a book
* `DELETE` : Delete a book

#Post Params
```
{
	"author": "Op Super John Doe Bilw",
	"name": "Implementation Golang",
	"category": "Knowledge"
}
```
