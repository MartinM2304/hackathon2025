package models

type Sound struct {
	Id     byte
	IpAddr string
}

func (e *Sound) DbData() DbData {
	return DbData{
		DataType: "sound",
		Value:    int(e.Id),
		IpAddr:   e.IpAddr,
	}
}

const (
	Bark byte = iota
	MeowMeow
	Opera
	DrumBeat
)

func soundToString(sound byte) string {
	switch sound {
	case 0:
		return "Bark"
	case 1:
		return "Meow Meow"
	case 2:
		return "Opera"
	case 3:
		return "Drum beat"
	}
	return ""
}
