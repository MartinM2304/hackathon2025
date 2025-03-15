package models

import "strconv"

type AggregatedData struct {
	Direction *AgreggatedDirection `json:"direction"`
	Emoji     *AgreggatedEmoji     `json:"emoji"`
	Sound     *AgreggatedSound     `json:"sound"`
}

func (a AggregatedData) ToString() string {
	directionStr := "No data"
	if a.Direction != nil {
		directionStr = directionToString(
			a.Direction.Id,
		) + "(" + strconv.Itoa(
			a.Direction.Count,
		) + ")"
	}

	emojiStr := "No data"
	if a.Emoji != nil {
		emojiStr = emojiToString(a.Emoji.Id) + "(" + strconv.Itoa(a.Emoji.Count) + ")"
	}

	soundStr := "No data"
	if a.Sound != nil {
		soundStr = soundToString(a.Sound.Id) + "(" + strconv.Itoa(a.Sound.Count) + ")"
	}

	return "Direction: " + directionStr + ", Emoji: " + emojiStr + ", Sound: " + soundStr
}
