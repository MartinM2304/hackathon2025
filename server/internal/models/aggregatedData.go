package models

type AggregatedData struct {
	Direction *byte `json:"direction"`
	Emoji     *byte `json:"emoji"`
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
