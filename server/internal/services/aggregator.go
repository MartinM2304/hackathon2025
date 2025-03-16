package services

import (
	"fmt"
	"math"
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
	sounds      []models.Sound
	soundsMutex sync.Mutex
)

var (
	aggregatedDirections     models.Queue[models.AgreggatedDirection]
	aggregatedDirectionMutex sync.Mutex
)

var (
	aggregatedEmojis      models.Queue[models.AgreggatedEmoji]
	aggregatedEmojisMutex sync.Mutex
)

var (
	aggregatedSounds      models.Queue[models.AgreggatedSound]
	aggregatedSoundsMutex sync.Mutex
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

	if maxCount == 0 {
		return
	}

	aggregatedDirectionMutex.Lock()
	aggregatedDirections.Enqueue(models.AgreggatedDirection{
		Id:    maxIdx,
		Count: maxCount,
	})
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

	if maxCount == 0 {
		return
	}

	aggregatedEmojisMutex.Lock()
	aggregatedEmojis.Enqueue(models.AgreggatedEmoji{
		Id:    maxIdx,
		Count: maxCount,
	})
	aggregatedEmojisMutex.Unlock()
}

func aggregateSounds() {
	soundsCounter := [4]int{0, 0, 0, 0}
	for _, sound := range sounds {
		soundsCounter[sound] += 1
	}

	maxIdx := models.Bark
	maxCount := soundsCounter[maxIdx]
	for i, count := range soundsCounter {
		if count > maxCount {
			maxCount = count
			maxIdx = models.Sound(i)
		}
	}

	if maxCount == 0 {
		return
	}

	aggregatedSoundsMutex.Lock()
	aggregatedSounds.Enqueue(models.AgreggatedSound{
		Id:    maxIdx,
		Count: maxCount,
	})
	aggregatedSoundsMutex.Unlock()
}

func Aggregate() {
	aggregateDirections()
	aggregateEmojis()
	aggregateSounds()

	fmt.Println("Aggregating")

	items := []models.DBser{}

	for _, emoji := range emojis {
		items = append(items, emoji)
	}

	for _, direction := range directions {
		items = append(items, &direction)
	}

	for _, sound := range sounds {
		items = append(items, &sound)
	}

	if len(items) != 0 {
		err := database.BatchInsertItems(items)
		if err != nil {
			fmt.Printf("Errror %v\n", err)
		}
	}

	directionsMutex.Lock()
	directions = []models.Direction{}
	directionsMutex.Unlock()

	emojisMutex.Lock()
	emojis = []models.Emoji{}
	emojisMutex.Unlock()

	soundsMutex.Lock()
	sounds = []models.Sound{}
	soundsMutex.Unlock()
}

func calculateDirectionEntropy(userVotes map[int][4]int) []models.Pair {
	var result []models.Pair

	for turn, votes := range userVotes {
		totalVotes := [4]float64{0, 0, 0, 0}
		total := 0.0

		for dir := 0; dir < 4; dir++ {
			totalVotes[dir] = float64(votes[dir])
			total += totalVotes[dir]
		}

		if total == 0 {
			result = append(result, models.Pair{Turn: turn, Entropy: 0})
			continue
		}

		entropy := 0.0
		for _, count := range totalVotes {
			if count > 0 {
				p := count / total
				entropy -= p * math.Log2(p)
			}
		}

		maxEntropy := math.Log2(4)
		normalizedEntropy := entropy / maxEntropy
		entropyPercent := int(math.Round(normalizedEntropy * 100))

		result = append(result, models.Pair{Turn: turn, Entropy: entropyPercent})
	}

	return result
}
