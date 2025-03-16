package models

type StatsJson struct {
	Enthropy  []EntropyPair `json:"entropy"`
	Emoji     []StatPair    `json:"emoji"`
	Direction []StatPair    `json:"direction"`
	Sound     []StatPair    `json:"sound"`
}
