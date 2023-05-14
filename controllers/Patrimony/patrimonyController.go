package controllers

import (
	"net/http"
	models "tab/models/Patrimony"
	"tab/service"

	"github.com/gin-gonic/gin"
)

func GetAllPatrimony(c *gin.Context) {
	var patrimony []models.Patrimony
	service.DB.Find(&patrimony)

	c.JSON(http.StatusOK, gin.H{"data": patrimony})
}

func GetByIdPatrimony(c *gin.Context) {
	var patrimony models.Patrimony

	if err := service.DB.Where("id = ?", c.Param("id")).First(&patrimony).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patrimony not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patrimony})
}

func CreatePatrimony(c *gin.Context) {
	var input models.CreatePatrimonyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patrimony := models.Patrimony{Name: input.Name, Model: input.Model}
	service.DB.Create(&patrimony)

	c.JSON(http.StatusOK, gin.H{"data": patrimony})
}
