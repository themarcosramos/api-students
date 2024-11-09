package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

  type Student struct{
	gorm.Model
	 Name   string `json:"name"`
	 CPF    string `json:"cpf"`
	 Email  string `json:"email"`
	 Age    int    `json:"age"`
	 Active bool   `json:"registration"`
  }

  func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
  }

  func AddStudent(student Student) error{
	 db:= Init()

	if result:= db.Create(&student); result.Error != nil {
		return result.Error
	}

	fmt.Println("Create student !")
	return nil
  }

  func GetStudents() ([]Student, error) {
	students := []Student{}

	db:= Init()

	err := db.Find(&students).Error
	return students, err


  }
