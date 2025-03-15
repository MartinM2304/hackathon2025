package services

import (
	"sync"

	"github.com/MartinM2304/hackathon2025/internal/models"
)

var (
	directions      []models.Direction
	directionsMutex sync.Mutex
)

var (
	aggregatedDirections     models.Queue[models.Direction]
	aggregatedDirectionMutex sync.Mutex
)

func Aggregate() error {
	directionsCounter := []int{0, 0, 0, 0}
	for _, direction := range directions {
		directionsCounter[direction] += 1
	}

	maxIdx := models.Up
	maxCount := directionsCounter[maxIdx]
	for i, count := range directionsCounter {
		if count > maxCount {
			maxCount = count
			maxIdx = models.Direction(i)
		}
	}

	directionsMutex.Lock()
	directions = []models.Direction{}
	directionsMutex.Unlock()

	if maxCount == 0 {
		return nil
	}

	aggregatedDirectionMutex.Lock()
	aggregatedDirections.Enqueue(maxIdx)
	aggregatedDirectionMutex.Unlock()

	return nil
}
