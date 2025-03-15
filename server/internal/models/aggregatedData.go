package models

type AggregatedData struct {
	Direction *byte `json:"direction"`
	Emoji     *byte `json:"emoji"`
	Sound     *byte `json:"sound"`
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

	soundStr := "No data"
	if a.Sound != nil {
		soundStr = soundToString(*a.Sound)
	}

	return "Direction: " + directionStr + ", Emoji: " + emojiStr + ", Sound: " + soundStr
}
