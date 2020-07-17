// controllers/cars.go

package controllers

import (
	"github.com/chillicoder/cars/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /cars
// Get all cars
func FindCars(c *gin.Context) {
	var cars []models.Car
	models.DB.Find(&cars)

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// POST /cars
// Create new car
func CreateCar(c *gin.Context) {
	// Validate input
	var input models.CreateCarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	// Create car
	car := models.Car{Description: input.Description, Location: input.Location}
	models.DB.Create(&car)

	c.JSON(http.StatusOK, gin.H{"data": car})
}

// GET /cars/:id
// Find a car
func FindCar(c *gin.Context) {
	var car models.Car

	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": car})
}

// PATCH /cars/:id
// Update a car
func UpdateCar(c *gin.Context) {
	var car models.Car
	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Validate input
	var input models.UpdateCarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&car).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": car})
}

// DELETE /cars/:id
// Delete a car
func DeleteCar(c *gin.Context) {
	//
	var car models.Car
	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&car)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
