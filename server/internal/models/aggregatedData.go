package models

type AggregatedData struct {
	Direction *byte `json:"direction,omitempty"`
	Emoji     *byte `json:"emoji,omitempty"`
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
