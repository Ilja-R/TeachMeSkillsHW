package repository

import (
	"context"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/models"
)

var (
	repositoryName string = "employees"
)

func (r *Repository) GetAllEmployees(ctx context.Context) (employees []models.Employee, err error) {
	if err = r.db.SelectContext(ctx, &employees, `
		SELECT id, name, email, age
		FROM employees
		ORDER BY id`); err != nil {
		return nil, r.translateError(err, repositoryName)
	}
	if employees == nil {
		employees = []models.Employee{}
	}

	return employees, nil
}
func (r *Repository) GetEmployeeByID(ctx context.Context, id int) (employee models.Employee, err error) {
	if err = r.db.GetContext(ctx, &employee, `
		SELECT id, name, email, age
		FROM employees
		WHERE id = $1`, id); err != nil {
		return models.Employee{}, r.translateErrorWithId(err, repositoryName, id)
	}

	return employee, nil
}

func (r *Repository) CreateEmployee(ctx context.Context, employee models.Employee) (err error) {
	_, err = r.db.ExecContext(ctx, `INSERT INTO employees (name, email, age)
					VALUES ($1, $2, $3)`,
		employee.Name,
		employee.Email,
		employee.Age)
	if err != nil {
		return r.translateError(err, repositoryName)
	}

	return nil
}

func (r *Repository) UpdateEmployeeByID(ctx context.Context, user models.Employee) (err error) {
	_, err = r.db.ExecContext(ctx, `
		UPDATE employees SET name = $1, 
		                    email = $2, 
		                    age = $3
		                WHERE id = $4`,
		user.Name,
		user.Email,
		user.Age,
		user.ID)
	if err != nil {
		return r.translateErrorWithId(err, repositoryName, user.ID)
	}

	return nil
}

func (r *Repository) DeleteEmployeeByID(ctx context.Context, id int) (err error) {
	_, err = r.db.ExecContext(ctx, `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return r.translateErrorWithId(err, repositoryName, id)
	}

	return nil
}
