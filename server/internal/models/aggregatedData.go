package models

type AggregatedData struct {
	Direction *Direction `json:"direction"`
	Emoji     *Emoji     `json:"emoji"`
}

func (a AggregatedData) ToString() string {
	directionStr := "No data"
	if a.Direction != nil {
		directionStr = a.Direction.toString()
	}

	emojiStr := "No data"
	if a.Emoji != nil {
		emojiStr = a.Emoji.toString()
	}

	return "Direction: " + directionStr + ", Emoji: " + emojiStr
}
