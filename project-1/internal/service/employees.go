package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/errs"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/log"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/models"
	"time"
)

var (
	defaultTTL = time.Minute * 5
	logger     = log.WithLayer("employee_service")
)

func (s *EmployeeService) GetAllEmployees(ctx context.Context) (users []models.Employee, err error) {
	users, err = s.repository.GetAllEmployees(ctx)
	if err != nil {
		return nil, s.addErrorInfo("GetAllEmployees", err)
	}

	return users, nil
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id int) (employee models.Employee, err error) {

	// Get from cache
	err = s.cache.Get(ctx, employee.CacheKey(), &employee)
	if err == nil {
		logger.Debug().Msg(fmt.Sprintf("got employer %s from cache", employee.Name))
		return employee, nil
	}

	// Get from repository
	employee, err = s.repository.GetEmployeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return models.Employee{}, s.addErrorInfo("GetEmployeeByID", err)
		}
		return models.Employee{}, err
	}

	// Update cache
	if err = s.cache.Set(ctx, employee.CacheKey(), employee, defaultTTL); err != nil {
		logger.Warn().Err(err).Msg("could not set employee cache")
	}

	return employee, nil
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, user models.Employee) (err error) {
	err = s.repository.CreateEmployee(ctx, user)
	if err != nil {
		return s.addErrorInfo("CreateEmployee", err)
	}
	return nil
}

func (s *EmployeeService) UpdateEmployeeByID(ctx context.Context, user models.Employee) (err error) {
	employee, err := s.repository.GetEmployeeByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return s.addErrorInfo("GetEmployeeByID", fmt.Errorf("%s: %w", err.Error(), errs.ErrEmployeeNotfound))
		}
		return err
	}

	err = s.repository.UpdateEmployeeByID(ctx, user)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return s.addErrorInfo("UpdateEmployeeByID", fmt.Errorf("%s: %w", err.Error(), errs.ErrEmployeeNotfound))
		}
		return s.addErrorInfo("UpdateEmployeeByID", err)
	}
	user.ID = employee.ID

	// Update cache (only in case of success)
	if err = s.cache.Set(ctx, user.CacheKey(), user, defaultTTL); err != nil {
		logger.Warn().Err(err).Msg("error during cache set")
	}

	return nil
}

func (s *EmployeeService) DeleteEmployeeByID(ctx context.Context, id int) (err error) {
	user, err := s.repository.GetEmployeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return s.addErrorInfo("DeleteEmployeeByID", fmt.Errorf("%s: %w", err.Error(), errs.ErrEmployeeNotfound))
		}
		return err
	}

	err = s.repository.DeleteEmployeeByID(ctx, id)
	if err != nil {
		return s.addErrorInfo("DeleteEmployeeByID", err)
	}

	_ = s.cache.Delete(ctx, user.CacheKey())
	return nil
}
