package main

import (
	"tab/controllers"
	"tab/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	service.ConnectDatabase()

	r.GET("/", controllers.FindPatrimony)

	r.Run()
}
