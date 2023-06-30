package models

import "errors"

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type Student struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password password `json:"-"`
}

type password struct {
	Plaintext *string
	Hash      []byte
}
