package models

type Emoji int

func (e Emoji) DbData() DbData {
	return DbData{
		DataType: "emoji",
		Value:    int(e),
	}
}

const (
	Smile Emoji = iota
	Sad
	Angry
	Wow
)

func emojiToString(emoji Emoji) string {
	switch emoji {
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
