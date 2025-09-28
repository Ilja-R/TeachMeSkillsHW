package controller

import (
	_ "github.com/Ilja-R/TeachMeSkillsHW/project-1/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (ctrl *Controller) RegisterEndpoints() {

	ctrl.router.GET("/ping", ctrl.Ping)
	ctrl.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ctrl.router.GET("/employees", ctrl.GetAllEmployees)
	ctrl.router.GET("/employees/:id", ctrl.GetEmployeeByID)
	ctrl.router.POST("/employees", ctrl.CreateEmployee)
	ctrl.router.PUT("/employees/:id", ctrl.UpdateEmployeeByID)
	ctrl.router.DELETE("/employees/:id", ctrl.DeleteEmployeeByID)
}

func (ctrl *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResponse{Message: "Server is up and running!"})
}

func (ctrl *Controller) RunServer(address string) error {
	// Регистрируем роуты
	ctrl.RegisterEndpoints()

	// Запускаем http-сервер
	if err := ctrl.router.Run(address); err != nil {
		return err
	}

	return nil
}
