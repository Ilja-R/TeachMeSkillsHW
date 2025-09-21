package service

import (
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/contracts"
)

type UsersService struct {
	repository contracts.UsersRepositoryI
}

func NewService(repository contracts.UsersRepositoryI) *UsersService {
	return &UsersService{
		repository: repository,
	}
}
