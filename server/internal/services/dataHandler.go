package services

import (
	"github.com/MartinM2304/hackathon2025/internal/database"
	"github.com/MartinM2304/hackathon2025/internal/models"
)

func RegisterDirectionVote(direction models.Direction) {
	directionsMutex.Lock()
	defer directionsMutex.Unlock()

	directions = append(directions, direction)
}

func RegisterEmojiVote(emoji models.Emoji) {
	emojisMutex.Lock()
	defer emojisMutex.Unlock()

	emojis = append(emojis, emoji)
}

func RegisterSoundVote(sound models.Sound) {
	soundsMutex.Lock()
	defer soundsMutex.Unlock()

	sounds = append(sounds, sound)
}

func GetAggregatedData() models.AggregatedData {
	aggregatedData := models.AggregatedData{}

	if !aggregatedDirections.IsEmpty() {
		aggregatedDirectionMutex.Lock()
		direction := aggregatedDirections.Dequeue()
		aggregatedDirectionMutex.Unlock()

		aggregatedData.Direction = &direction
	} else {
		aggregatedData.Direction = nil
	}

	if !aggregatedEmojis.IsEmpty() {
		aggregatedEmojisMutex.Lock()
		emoji := aggregatedEmojis.Dequeue()
		aggregatedEmojisMutex.Unlock()

		aggregatedData.Emoji = &emoji
	} else {
		aggregatedData.Emoji = nil
	}

	if !aggregatedSounds.IsEmpty() {
		aggregatedSoundsMutex.Lock()
		sound := aggregatedSounds.Dequeue()
		aggregatedSoundsMutex.Unlock()

		aggregatedData.Sound = &sound
	} else {
		aggregatedData.Sound = nil
	}

	sendNotification(aggregatedData)

	return aggregatedData
}

func CalculateEnthropy() ([]models.Pair, error) {
	err, directionsData := database.GetAllItemsForType("direction")
	if err != nil {
		return nil, err
	}
	enthropy := calculateDirectionEntropy(directionsData)
	return enthropy, nil
}
