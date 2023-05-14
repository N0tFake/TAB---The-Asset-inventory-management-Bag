package controllers

import (
	"net/http"
	models "tab/models/Patrimony"
	"tab/service"

	"github.com/gin-gonic/gin"
)

func DeletePatrimony(c *gin.Context) {
	var patrimony models.Patrimony
	if err := service.DB.Where("id = ?", c.Param("id")).First(&patrimony).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	service.DB.Delete(&patrimony)
	c.JSON(http.StatusOK, gin.H{"delete": true})
}
