package main

import (
    "log"

    "github.com/gin-contrib/cors"
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

    r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"},
    AllowMethods:     []string{"GET", "POST", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    AllowCredentials: true,
}))


    routes.InitSensorRoutes(r)

    r.Run(":8080")
}
