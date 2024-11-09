package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/themarcosramos/api-students/db"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes

  e.GET("/students", getStudents)
  e.POST("/students", createStudent)

  e.GET("/students/:id", getStudent)
  e.PUT("/students/:id", updateStudent)

  e.DELETE("/students/:id", deleteStudent)
 
  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getStudents(c echo.Context) error {
  return c.String(http.StatusOK, "List all students!")
}

func createStudent(c echo.Context) error {
	db.AddStudent()
	return c.String(http.StatusOK, "Create student!")
}
  
func getStudent(c echo.Context) error {
	id:= c.Param("id")
	getSud := fmt.Sprintf("Get %s student",id)
	return c.String(http.StatusOK, getSud)
}
  
func updateStudent(c echo.Context) error {
	id:= c.Param("id")
	getSud := fmt.Sprintf("Update %s student",id)
	return c.String(http.StatusOK, getSud)
}
  
func deleteStudent(c echo.Context) error {
	id:= c.Param("id")
	getSud := fmt.Sprintf("Delete %s student",id)
	return c.String(http.StatusOK, getSud)
}
  