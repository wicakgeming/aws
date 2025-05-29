package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/wicakgeming/aws/config"
	"github.com/wicakgeming/aws/models"

)

func InitSensorRoutes(r *gin.Engine) {
	r.POST("/sensor", func(c *gin.Context) {
		var input models.SensorData
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		input.GPSDateTime = input.GPSDateTime.UTC() // Normalize datetime
		result := config.DB.Create(&input)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, input)
	})
}
