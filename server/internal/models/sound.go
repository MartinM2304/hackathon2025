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
	Hello
	Bye
)

func soundToString(sound Sound) string {
	switch sound {
	case Bark:
		return "Bark"
	case MeowMeow:
		return "Meow Meow"
	case Hello:
		return "Hello"
	case Bye:
		return "Bye"
	}
	return ""
}
