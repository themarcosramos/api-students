package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/themarcosramos/api-students/schemas"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
	student := schemas.Student{}
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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	receivedStudent := schemas.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	updatingStudent, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Students not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	student := updateStudentInfo(receivedStudent, updatingStudent)

	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save  student")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) deleteStudent(c echo.Context) error {
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

	if err := api.DB.DeleteStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete  student")
	}

	return c.JSON(http.StatusOK, student)
}

func updateStudentInfo(receivedStudent, student schemas.Student) schemas.Student {

	if receivedStudent.Name != "" {
		student.Name = receivedStudent.Name
	}

	if receivedStudent.CPF != "" {
		student.CPF = receivedStudent.CPF
	}

	if receivedStudent.Email != "" {
		student.Email = receivedStudent.Email
	}

	if receivedStudent.Age > 0 {
		student.Age = receivedStudent.Age
	}

	if receivedStudent.Active != student.Active {
		student.Active = receivedStudent.Active
	}

	return student
}
