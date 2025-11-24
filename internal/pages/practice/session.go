package practice

import (
	"math/rand"
)

type Glyph struct {
	Char  string
	State GlyphState
}

type GlyphState int

const (
	Pending GlyphState = iota
	Correct
	Wrong
)

var (
	lowerCaseLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func createSessionText() []Glyph {
	var sessionText []Glyph

	for _ = range 5 {
		var text = Glyph{
			Char:  lowerCaseLetters[rand.Intn(25)],
			State: 0,
		}
		sessionText = append(sessionText, text)
	}

	return sessionText
}

// Returns the current session with empty input and Glyph states back to pending
func RestartSessionText(curr PracticeModel) PracticeModel {
	var g []Glyph

	for _, t := range curr.TargetText {
		var c Glyph

		c = Glyph{
			Char:  t.Char,
			State: 0,
		}

		g = append(g, c)
	}

	return PracticeModel{
		TargetText: g,
	}
}

// Compares Input vs Target, returns the changed Glyph state (correct or wrong)
func evaluateInput(i string, g Glyph) Glyph {
	if i == g.Char {
		g.State = 1
	} else {
		g.State = 2
	}
	return g
}
