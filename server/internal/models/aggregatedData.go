package models

type AggregatedData struct {
	Direction *byte `json:"direction"`
	Emoji     *byte `json:"emoji"`
}

func (a AggregatedData) ToString() string {
	directionStr := "No data"
	if a.Direction != nil {
		directionStr = directionToString(*a.Direction)
	}

	emojiStr := "No data"
	if a.Emoji != nil {
		emojiStr = emojiToString(*a.Emoji)
	}

	return "Direction: " + directionStr + ", Emoji: " + emojiStr
}
