package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wicakgeming/aws/controller"
)

func InitSensorRoutes(r *gin.Engine) {
	r.POST("/sensor", controllers.CreateSensorData)
}
