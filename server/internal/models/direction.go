package models

type Direction int

func (e Direction) DbData() DbData {
	return DbData{
		DataType: "direction",
		Value:    int(e),
	}
}

const (
	Up Direction = iota
	Down
	RotateLeft
	RotateRight
)

func directionToString(direction Direction) string {
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
