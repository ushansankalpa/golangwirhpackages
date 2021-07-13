package controller

import (
	"net/http"

	"grom_with_echo/model"
	"grom_with_echo/storage"

	"github.com/labstack/echo/v4"
)

func GetStudent(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Students, error) {
	db := storage.GetDBInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
