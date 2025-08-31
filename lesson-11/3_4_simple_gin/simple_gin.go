package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Division struct {
	A int `json:"a" binding:"required"`
	B int `json:"b" binding:"required"`
}

func helloHandler(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is empty"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello %s!", name)})
}

func divisionHandler(c *gin.Context) {
	var division Division
	if err := c.BindJSON(&division); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if division.B == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "division by zero"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": division.A / division.B})
}

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.POST("/divide", divisionHandler)
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
