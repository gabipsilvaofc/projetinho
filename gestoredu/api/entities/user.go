package entities

import "github.com/google/uuid"

type User struct {
	ID             string `json:"id"`
    Name           string `json:"name"`
    Sobrenome      string `json:"sobrenome"`
    Telefone       int    `json:"telefone"`
    Email          string `json:"email"`
    Sexo           string `json:"sexo"`
    CPF            int    `json:"cpf"`
    Endereco       string `json:"endereco"`
    DataNascimento int    `json:"dataNascimento"`
    Matricula      string `json:"matricula"`
    Tipo           string `json:"tipo"`
}

func NewUser() *User {
	user := User{
		ID: uuid.New().String(),
	}
	return &user
}