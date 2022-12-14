package service

import (
	"github.com/Dimiyrg/GoProject4"
	"github.com/Dimiyrg/GoProject4/pkg/repository"
)

type Authorization interface {
	CreateUser(user GoProject4.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
