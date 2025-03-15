package models

type AggregatedData struct {
	Direction *Direction `json:"direction,omitempty"`
	Emoji     *Emoji     `json:"emoji,omitempty"`
}

func (a AggregatedData) ToString() string {
	directionStr := "nil"
	if a.Direction != nil {
		directionStr = a.Direction.toString()
	}

	emojiStr := "nil"
	if a.Emoji != nil {
		emojiStr = a.Emoji.toString()
	}

	return "Direction: " + directionStr + ", Emoji: " + emojiStr
}
