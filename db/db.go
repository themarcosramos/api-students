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

  func AddStudent(student Student) {
	 db:= Init()


	if result:= db.Create(&student); result.Error != nil {
		fmt.Println("Error to create student")
	}

	fmt.Println("Create student !")
  }
