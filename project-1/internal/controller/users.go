package controller

import (
	"errors"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/errs"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllUsers
// @Summary Get all users
// @Description Get all users or throw an error
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} CommonError
// @Router /users [get]
func (ctrl *Controller) GetAllUsers(c *gin.Context) {
	products, err := ctrl.service.GetAllUsers(c)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetUserByID
// @Summary Get user by ID
// @Description Get user by ID or throw an error
// @Tags Users
// @Produce json
// @Param id path int true "id of user"
// @Success 200 {object} models.User
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (ctrl *Controller) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	product, err := ctrl.service.GetUserByID(c, id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateUser
// @Summary Add new user
// @Description Add new user or throw an error
// @Tags Users
// @Consume json
// @Produce json
// @Param request_body body dto.UserCreateRequest true "new user info"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users [post]
func (ctrl *Controller) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if user.Name == "" || user.Email == "" || user.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err := ctrl.service.CreateUser(c, user); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully"})
}

// UpdateUserByID
// @Summary Update user by ID
// @Description Update user by ID or throw an error
// @Tags Users
// @Consume json
// @Produce json
// @Param id path int true "user id"
// @Param request_body body dto.UserCreateRequest true "user info"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [put]
func (ctrl *Controller) UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	var user models.User
	if err = c.ShouldBindJSON(&user); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if user.Name == "" || user.Email == "" || user.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	user.ID = id

	if err = ctrl.service.UpdateUserByID(c, user); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "User updated successfully"})
}

// DeleteUserByID
// @Summary Delete user by ID
// @Description Delete user by ID or throw an error
// @Tags Users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [delete]
func (ctrl *Controller) DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	if err = ctrl.service.DeleteUserByID(c, id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "User deleted successfully"})
}
