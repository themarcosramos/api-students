package schemas

import (
	"time"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type StudentResponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Active    bool      `json:"active"`
}

func NewResponse(students []Student) []StudentResponse {
	studentsResponse := []StudentResponse{}

	for _, student := range students {
		studentResponse := StudentResponse{
			ID:        int(student.ID),
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			Name:      student.Name,
			Email:     student.Email,
			Age:       student.Age,
			CPF:       student.CPF,
			Active:    student.Active,
		}
		studentsResponse = append(studentsResponse, studentResponse)
	}
	return studentsResponse
}
