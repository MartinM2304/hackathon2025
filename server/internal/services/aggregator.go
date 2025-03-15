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
	emojis      []models.Emoji
	emojisMutex sync.Mutex
)

var (
	aggregatedDirections     models.Queue[models.Direction]
	aggregatedDirectionMutex sync.Mutex
)

var (
	aggregatedEmojis      models.Queue[models.Emoji]
	aggregatedEmojisMutex sync.Mutex
)

func aggregateDirections() {
	directionsCounter := [4]int{0, 0, 0, 0}
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
		return
	}

	aggregatedDirectionMutex.Lock()
	aggregatedDirections.Enqueue(maxIdx)
	aggregatedDirectionMutex.Unlock()
}

func aggregateEmojis() {
	emojisCounter := [4]int{0, 0, 0, 0}
	for _, emoji := range emojis {
		emojisCounter[emoji] += 1
	}

	maxIdx := models.Smile
	maxCount := emojisCounter[maxIdx]
	for i, count := range emojisCounter {
		if count > maxCount {
			maxCount = count
			maxIdx = models.Emoji(i)
		}
	}

	emojisMutex.Lock()
	emojis = []models.Emoji{}
	emojisMutex.Unlock()

	if maxCount == 0 {
		return
	}

	aggregatedEmojisMutex.Lock()
	aggregatedEmojis.Enqueue(maxIdx)
	aggregatedEmojisMutex.Unlock()
}

func Aggregate() {
	aggregateDirections()
	aggregateEmojis()
}
