package models

type Emoji byte

const (
	Smile Emoji = iota
	Sad
	Angry
	Wow
)

func (e Emoji) toString() string {
	switch e {
	case Smile:
		return "Smile"
	case Sad:
		return "Sad"
	case Angry:
		return "Angry"
	case Wow:
		return "Wow"
	default:
		return "Unknown"
	}
}
