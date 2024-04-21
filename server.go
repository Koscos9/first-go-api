package main

import (
	"nii-git/first-go-api/api"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	maxDBRetryCount = 3
	sleepTime       = 10
)

func main() {
	e := echo.New()

	// db接続
	db, err := connectDB()
	if err != nil {
		e.Logger.Fatal(err.Error())
		return
	}

	// api
	api := api.Api{
		Db: db,
	}

	// router
	e.GET("/", api.HelloWorld())
	e.GET("/test", api.GetTestData())
	e.POST("/test", api.PostTestData())

	e.Logger.Fatal(e.Start(":1323"))
}

func connectDB() (db *gorm.DB, err error) {
	for r := 1; r <= maxDBRetryCount; r++ {
		db, err = gorm.Open(mysql.Open("root:password@tcp(mysql:3306)/go-sample?charset=utf8mb4&parseTime=True&loc=Local"))
		if err == nil {
			return db, nil
		}
		time.Sleep(sleepTime * time.Second)
	}
	return nil, err
}
