package tools

import (
	"math/rand"
	"time"
)

func GenerateRandomAsciiString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyz1234567890"
	aray := letters[0:length]
	str := ""
	for _ = range aray {
		i := rand.Intn(len(letters))
		s := letters[i : i+1]
		str = str + s
	}
	return str
}
