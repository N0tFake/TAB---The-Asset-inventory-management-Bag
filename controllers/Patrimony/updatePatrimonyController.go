package controllers

import (
	"net/http"
	models "tab/models/Patrimony"
	"tab/service"

	"github.com/gin-gonic/gin"
)

func UpdatePatrimony(c *gin.Context) {
	var patrimony models.Patrimony

	if err := service.DB.Where("id = ?", c.Param("id")).First(&patrimony).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdatePatrimonyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.DB.Model(&patrimony).Updates(map[string]interface{}{"Name": input.Name, "model": input.Model})
	c.JSON(http.StatusOK, gin.H{"data": patrimony})
}
