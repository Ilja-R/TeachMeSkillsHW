package contracts

import (
	"context"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/models"
)

type EmployeeServiceI interface {
	GetAllEmployees(ctx context.Context) (employees []models.Employee, err error)
	GetEmployeeByID(ctx context.Context, id int) (employee models.Employee, err error)
	CreateEmployee(ctx context.Context, employee models.Employee) (err error)
	UpdateEmployeeByID(ctx context.Context, employee models.Employee) (err error)
	DeleteEmployeeByID(ctx context.Context, id int) (err error)
}
