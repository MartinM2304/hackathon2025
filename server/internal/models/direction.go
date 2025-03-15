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

func directionToString(direction byte) string {
	switch direction {
	case 0:
		return "Up"
	case 1:
		return "Down"
	case 2:
		return "RotateLeft"
	case 3:
		return "RotateRight"
	}
	return ""
}
