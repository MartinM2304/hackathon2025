package services

import "github.com/MartinM2304/hackathon2025/internal/models"

var NotificationChannel = make(chan string)

func sendNotification(aggregatedData models.AggregatedData) {
	NotificationChannel <- aggregatedData.ToString()
}
