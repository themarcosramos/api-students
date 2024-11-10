package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool  `json:"active"` // força que a entrada true/false
}

func erroParamRequest(param, typ string) error {
	return fmt.Errorf("param '%s' of type '%s'is required", param, typ)
}

func (s *StudentRequest) Validate() error {

	if s.Name == "" {
		return erroParamRequest("name", "string")
	}
	if s.CPF == "" {
		return erroParamRequest("cpf", "string")
	}
	if s.Email == "" {
		return erroParamRequest("email", "string")
	}
	if s.Age == 0 {
		return erroParamRequest("age", "int")
	}
	if s.Active == nil {
		return erroParamRequest("active", "bool")
	}

	return nil
}
