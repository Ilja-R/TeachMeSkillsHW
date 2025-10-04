package controller

import (
	"errors"
	"fmt"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/contracts"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	router  *gin.Engine
	service contracts.EmployeeServiceI
}

func NewController(service contracts.EmployeeServiceI) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}
}

func (ctrl *Controller) handleError(c *gin.Context, err error) {
	logger.Error().Err(err).Msg(fmt.Sprintf("Error: %v", c.Request.UserAgent()))
	switch {
	case errors.Is(err, errs.ErrEmployeeNotfound) || errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidEmployeeID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
