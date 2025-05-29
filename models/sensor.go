package models

import "time"

type SensorData struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	WaterLevel   float64   `json:"waterLevel"`
	WindSpeed    float64   `json:"windSpeed"`
	AirTemp      float64   `json:"airTemp"`
	AirHumidity  float64   `json:"airHumidity"`
	AirPressure  float64   `json:"airPressure"`
	WaterTemp    float64   `json:"waterTemp"`
	CurahHujan   float64   `json:"curahHujan"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	GPSDateTime  time.Time `json:"gpsDateTime"`
	CreatedAt    time.Time
}
