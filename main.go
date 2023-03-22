package main

import (
	"tab/controllers"
	"tab/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	service.ConnectDatabase()

	r.GET("/", controllers.GetAllPatrimony)
	r.POST("/patrimony", controllers.CreatePatrimony)
	r.GET("/patrimony/:id", controllers.GetByIdPatrimony)

	r.Run()
}
