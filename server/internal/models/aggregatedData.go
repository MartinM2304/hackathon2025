package models

type AggregatedData struct {
	Direction *Direction `json:"direction,omitempty"`
	Emoji     *Emoji     `json:"emoji,omitempty"`
}

func (a *AggregatedData) ToString() string {
	return "Direction: " + a.Direction.toString() + ", Emoji: " + a.Emoji.toString()
}
