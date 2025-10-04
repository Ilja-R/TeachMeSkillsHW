package service

import (
	"fmt"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/contracts"
)

type EmployeeService struct {
	repository contracts.EmployeeRepositoryI
	cache      contracts.CacheI
}

func NewService(repository contracts.EmployeeRepositoryI, cache contracts.CacheI) *EmployeeService {
	return &EmployeeService{
		repository: repository,
		cache:      cache,
	}
}

func (s *EmployeeService) addErrorInfo(methodName string, err error) error {
	return fmt.Errorf("employer service %s: %w", methodName, err)
}
