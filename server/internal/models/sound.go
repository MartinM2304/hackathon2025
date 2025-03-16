package models

type Sound int

func (e Sound) DbData() DbData {
	return DbData{
		DataType: "sound",
		Value:    int(e),
	}
}

const (
	Bark Sound = iota
	MeowMeow
	Opera
	DrumBeat
)

func soundToString(sound Sound) string {
	switch sound {
	case Bark:
		return "Bark"
	case MeowMeow:
		return "Meow Meow"
	case Opera:
		return "Opera"
	case DrumBeat:
		return "Drum beat"
	}
	return ""
}
