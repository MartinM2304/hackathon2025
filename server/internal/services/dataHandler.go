package services

import (
	"errors"

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

func GetAggregatedData() (error, models.AggregatedData) {
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

	if aggregatedData.Direction == nil && aggregatedData.Emoji == nil {
		return errors.New("No data available"), aggregatedData
	}

	sendNotification(aggregatedData)

	return nil, aggregatedData
}
