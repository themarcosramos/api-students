package api

import (
	"fmt"
	"errors"
	"net/http"
	"strconv"

	"github.com/themarcosramos/api-students/db"
	"gorm.io/gorm"


	"github.com/labstack/echo/v4"
)

// Handler
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student!")
	}

	return c.String(http.StatusOK, "Create student!")
}

func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	student, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		 return c.String(http.StatusNotFound, "Students not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	
	return c.JSON(http.StatusOK, student)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	getSud := fmt.Sprintf("Update %s student", id)
	return c.String(http.StatusOK, getSud)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	getSud := fmt.Sprintf("Delete %s student", id)
	return c.String(http.StatusOK, getSud)
}
