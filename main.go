package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/wicakgeming/aws/config"
	"github.com/wicakgeming/aws/models"
	"github.com/wicakgeming/aws/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.SensorData{})

	r := gin.Default()
	routes.InitSensorRoutes(r)

	r.Run(":8080")
}
