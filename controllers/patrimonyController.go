package controllers

import (
	"net/http"
	"tab/models"
	"tab/service"

	"github.com/gin-gonic/gin"
)

func FindPatrimony(c *gin.Context) {
	var patrimony []models.Patrimony
	service.DB.Find(&patrimony)

	c.JSON(http.StatusOK, gin.H{"data": patrimony})
}
