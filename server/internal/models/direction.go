package models

type Direction struct {
	Id     byte
	IpAddr string
}

func (e *Direction) DbData() DbData {
	return DbData{
		DataType: "direction",
		Value:    int(e.Id),
	}
}

const (
	Up byte = iota
	Down
	RotateLeft
	RotateRight
)

func directionToString(direction byte) string {
	switch direction {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case RotateLeft:
		return "RotateLeft"
	case RotateRight:
		return "RotateRight"
	}
	return ""
}
