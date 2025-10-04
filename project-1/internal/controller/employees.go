package controller

import (
	"errors"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/errs"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/log"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	logger = log.WithLayer("controller")
)

// GetAllEmployees
// @Summary Get all employees
// @Description Get all employees or throw an error
// @Tags Employees
// @Produce json
// @Success 200 {array} models.Employee
// @Failure 500 {object} CommonError
// @Router /employees [get]
func (ctrl *Controller) GetAllEmployees(c *gin.Context) {
	employees, err := ctrl.service.GetAllEmployees(c)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, employees)
}

// GetEmployeeByID
// @Summary Get employee by ID
// @Description Get employee by ID or throw an error
// @Tags Employees
// @Produce json
// @Param id path int true "id of employee"
// @Success 200 {object} models.Employee
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /employees/{id} [get]
func (ctrl *Controller) GetEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidEmployeeID)
		return
	}

	employee, err := ctrl.service.GetEmployeeByID(c, id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

// CreateEmployee
// @Summary Add new employee
// @Description Add new employee or throw an error
// @Tags Employees
// @Consume json
// @Produce json
// @Param request_body body dto.EmployeeCreateRequest true "new employee info"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /employees [post]
func (ctrl *Controller) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if employee.Name == "" || employee.Email == "" || employee.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err := ctrl.service.CreateEmployee(c, employee); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "Employee created successfully"})
}

// UpdateEmployeeByID
// @Summary Update employee by ID
// @Description Update employee by ID or throw an error
// @Tags Employees
// @Consume json
// @Produce json
// @Param id path int true "employee id"
// @Param request_body body dto.EmployeeCreateRequest true "employee info"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /employees/{id} [put]
func (ctrl *Controller) UpdateEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidEmployeeID)
		return
	}

	var employee models.Employee
	if err = c.ShouldBindJSON(&employee); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if employee.Name == "" || employee.Email == "" || employee.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	employee.ID = id

	if err = ctrl.service.UpdateEmployeeByID(c, employee); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Employee updated successfully"})
}

// DeleteEmployeeByID
// @Summary Delete employee by ID
// @Description Delete employee by ID or throw an error
// @Tags Employees
// @Produce json
// @Param id path int true "employee id"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /employee/{id} [delete]
func (ctrl *Controller) DeleteEmployeeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidEmployeeID)
		return
	}

	if err = ctrl.service.DeleteEmployeeByID(c, id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Employee deleted successfully"})
}
