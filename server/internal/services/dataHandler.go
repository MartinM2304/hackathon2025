package services

import (
	"errors"

	"github.com/MartinM2304/hackathon2025/internal/models"
)

func RegisterDirectionVote(direction models.Direction) {
	mutex.Lock()
	defer mutex.Unlock()

	directions = append(directions, direction)
}

func GetDirections() []models.Direction {
	return directions
}

func GetAggregatedDirection() (error, models.Direction) {
	if aggregatedDirection.GetLength() == 0 {
		return errors.New("No directions"), 0
	}

	aggregatedDirectionMutex.Lock()
	direction := aggregatedDirection.Dequeue()
	aggregatedDirectionMutex.Lock()
	return nil, direction
}
