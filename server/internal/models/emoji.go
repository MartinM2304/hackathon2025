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
