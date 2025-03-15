package models

type Direction struct {
	Id     byte
	IpAddr string
}

func (e *Direction) DbData() DbData {
	return DbData{
		DataType: "direction",
		Value:    int(e.Id),
		IpAddr:   e.IpAddr,
	}
}

const (
	Up byte = iota
	Down
	RotateLeft
	RotateRight
)

func (d Direction) toString() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case RotateLeft:
		return "RotateLeft"
	case RotateRight:
		return "RotateRight"
	default:
		return "Unknown"
	}
}
