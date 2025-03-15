package models

type AggregatedData struct {
	Direction *Direction `json:"direction,omitempty"`
	Emoji     *Emoji     `json:"emoji,omitempty"`
}
