package models

type Direction byte

const (
	Up Direction = iota
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
