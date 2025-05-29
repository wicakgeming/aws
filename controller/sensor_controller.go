package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wicakgeming/aws/config"
	"github.com/wicakgeming/aws/models"
)

func CreateSensorData(c *gin.Context) {
	var input models.SensorData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.GPSDateTime = input.GPSDateTime.UTC()

	result := config.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

func GetAllSensorData(c *gin.Context) {
	var sensors []models.SensorData

	if err := config.DB.Find(&sensors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sensors)
}

func DeleteSensorData(c *gin.Context) {
	id := c.Param("id")

	var sensor models.SensorData
	if err := config.DB.First(&sensor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	if err := config.DB.Delete(&sensor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
}

