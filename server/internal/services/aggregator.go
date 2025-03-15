package services

import (
	"sync"

	"github.com/MartinM2304/hackathon2025/internal/database"
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
	aggregatedDirections     models.Queue[byte]
	aggregatedDirectionMutex sync.Mutex
)

var (
	aggregatedEmojis      models.Queue[byte]
	aggregatedEmojisMutex sync.Mutex
)

func aggregateDirections() {
	directionsCounter := [4]int{0, 0, 0, 0}
	for _, direction := range directions {
		directionsCounter[direction.Id] += 1
	}

	maxIdx := models.Up
	maxCount := directionsCounter[maxIdx]
	for i, count := range directionsCounter {
		if count > maxCount {
			maxCount = count
			maxIdx = byte(i)
		}
	}

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
		emojisCounter[emoji.Id] += 1
	}

	maxIdx := models.Smile
	maxCount := emojisCounter[maxIdx]
	for i, count := range emojisCounter {
		if count > maxCount {
			maxCount = count
			maxIdx = byte(i)
		}
	}

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

	items := []models.DBser{}

	for _, emoji := range emojis {
		items = append(items, emoji)
	}

	for _, direction := range directions {
		items = append(items, &direction)
	}

	database.BatchInsertItems(items)

	directionsMutex.Lock()
	directions = []models.Direction{}
	directionsMutex.Unlock()

	emojisMutex.Lock()
	emojis = []models.Emoji{}
	emojisMutex.Unlock()
}
