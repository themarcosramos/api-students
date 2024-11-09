package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

  type Student struct{
	gorm.Model
	 Name string
	 CPF string
	 Email string
	 Age int
	 Active bool
  }

  func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
  }

  func AddStudent() {
	 db:= Init()

	 student := Student {
		Name:"Marcos Ramos",
		CPF :"012345678911",
		Email:"teste@mail.com",
		Age:30,
		Active:true,
	 }

	//  result:= db.Create(&student)
	//  if result.Error != nil{
	// 	fmt.Println("Error to create student")
	//  }

	if result:= db.Create(&student); result.Error != nil {
		fmt.Println("Error to create student")
	}

	fmt.Println("Create student !")
  }
