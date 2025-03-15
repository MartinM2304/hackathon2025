package models

type Emoji struct {
	Id     byte
	IpAddr string
}

func (e Emoji) DbData() DbData {
	return DbData{
		DataType: "emoji",
		Value:    int(e.Id),
		IpAddr:   e.IpAddr,
	}
}

const (
	Smile byte = iota
	Sad
	Angry
	Wow
)

func emojiToString(emoji byte) string {
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
