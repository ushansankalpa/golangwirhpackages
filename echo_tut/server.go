package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"Type"`
}

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the website")
}

func main() {
	fmt.Println("Hello Ushan Sankalpa")

	e := echo.New()

	e.GET("/", yallo)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCats)

	e.Start(":8000")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your cat name is: %s\t and his type is: %s", catName, catType))
	} else if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"naem": catName,
			"type": catType,
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "you need let us know if you want json or string",
		})
	}

}

func addCats(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed reading the request body for addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed reading the request body for addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your cat: %#v", cat)
	return c.String(http.StatusOK, "We got your cat ")
}
