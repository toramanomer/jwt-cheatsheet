package user

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type User struct {
	Username string
}

func NewUser(username string) *User {
	return &User{Username: username}
}
