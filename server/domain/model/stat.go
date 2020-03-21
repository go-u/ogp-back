package model

import "time"

type Stat struct {
	Date        time.Time
	FQDN        string
	Host        string
	Count       uint
	Title       string
	Description string
	Image       string
	Type        string
	Lang        string
}
