package user

import "errors"

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

var (
	ErrorEmailTaken    = errors.New("e-mail já cadastrado")
	ErrorUsernameTaken = errors.New("username já cadastrado")
	ErrorEmptyUsername = errors.New("username obrigatório")
)
