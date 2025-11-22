package practice

import "math/rand"

var (
	lowerCaseLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func createSessionText() []string {
	var sessionText []string

	for _ = range 100 {
		sessionText = append(sessionText, lowerCaseLetters[rand.Intn(25)])
	}

	return sessionText
}
