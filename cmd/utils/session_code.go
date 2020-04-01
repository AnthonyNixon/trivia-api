package utils

import "math/rand"

var letters = []rune("ABCDEFGHKMNPRSTWXYZ")

func NewSessionCode() (code string) {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
