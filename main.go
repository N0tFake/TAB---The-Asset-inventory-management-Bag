package main

import (
	"log"
	controllers "tab/controllers/Patrimony"
	"tab/initializers"
	"tab/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("Error: ", err)
	}

	service.ConnectDatabase(&config)

	r.GET("/", controllers.GetAllPatrimony)
	// r.POST("/patrimony", controllers.CreatePatrimony)
	// r.GET("/patrimony/:id", controllers.GetByIdPatrimony)
	// r.PATCH("/patrimony/update/:id", controllers.UpdatePatrimony)
	// r.DELETE("/patrimony/delete/:id", controllers.DeletePatrimony)

	err = r.Run()
	if err != nil {
		return
	}
}
