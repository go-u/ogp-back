package model

import "time"

type User struct {
	ID        uint64
	UID       string
	Name      string
	CreatedAt time.Time
}
