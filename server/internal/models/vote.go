package models

import "time"

type Vote struct {
	ID        int
	Type      string
	Value     int
	Turn      int
	Timestamp time.Time
}
