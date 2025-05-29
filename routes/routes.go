package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wicakgeming/aws/controllers"
)

func InitSensorRoutes(r *gin.Engine) {
	r.POST("/sensor", controllers.CreateSensorData)
}
