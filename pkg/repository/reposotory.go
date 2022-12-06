package repository

import (
	"github.com/Dimiyrg/GoProject4"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user GoProject4.User) (int, error)
	GetUser(email, password string) (GoProject4.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
