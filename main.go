package main

import (
	"github.com/chillicoder/cars/controllers"
	"github.com/chillicoder/cars/models"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/cars", controllers.FindCars)
	r.POST("/cars", controllers.CreateCar)
	r.GET("/cars/:id", controllers.FindCar)
	r.PATCH("/cars/:id", controllers.UpdateCar)
	r.DELETE("/cars/:id", controllers.DeleteCar)

	r.Run()
}
