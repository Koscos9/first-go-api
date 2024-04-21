package api

import (
	"fmt"
	"net/http"
	"nii-git/first-go-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Api struct {
	Db *gorm.DB
}

func (a *Api) HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	}
}

func (a *Api) GetTestData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var test models.Test
		a.Db.Last(&test)
		return c.JSON(http.StatusOK, test)
	}
}

func (a *Api) PostTestData() echo.HandlerFunc {
	return func(c echo.Context) error {
		test := new(models.Test)
		err := c.Bind(&test)
		if err != nil {
			fmt.Println(err.Error())
			return c.String(http.StatusInternalServerError, "ng")
		}
		result := a.Db.Create(&test)
		if result.Error != nil {
			fmt.Println(err.Error())
			return c.String(http.StatusInternalServerError, "ng")
		}

		return c.String(http.StatusOK, "ok")
	}
}
