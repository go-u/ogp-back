package model

import "time"

type Bookmark struct {
	UserID    uint64
	FQDN      string
	CreatedAt time.Time
}
