package practice

import (
	"math/rand"
	"time"
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

type StopWatch struct {
	start   time.Time
	running bool
	elapsed time.Duration
}

func (sw *StopWatch) Start() {
	if sw.running {
		return
	}

	sw.start = time.Now()
	sw.running = true
}

func (sw *StopWatch) Stop() {
	if !sw.running {
		return
	}

	sw.elapsed = time.Since(sw.start)
	sw.running = false
}

func (sw *StopWatch) Reset() {
	sw.start = time.Time{}
	sw.elapsed = 0
	sw.running = false
}

func (sw *StopWatch) Elapsed() time.Duration {
	if sw.running {
		return sw.elapsed + time.Since(sw.start)
	}

	return sw.elapsed
}

var (
	lowerCaseLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func createSessionText() []Glyph {
	var sessionText []Glyph

	for _ = range 50 {
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

func evaluateRawWPM(input []string, t time.Duration) int {
	if t <= 0 {
		return 0
	}
	chars := float64(len(input))
	words := chars / 5.0
	mins := t.Minutes()

	if mins == 0 {
		return 0
	}

	return int(words / mins)
}

func evaluateWPM(input []Glyph, t time.Duration) int {
	var c float64

	for _, g := range input {
		if g.State == 1 {
			c++
		}
	}

	chars := float64(len(input))
	acc := c / chars
	words := chars / 5.0
	mins := t.Minutes()

	if mins == 0 {
		return 0
	}

	return int((words * acc) / mins)
}

func evaluateAccuracy(target []Glyph, totalInput []string) float64 {
	var c float64

	for _, g := range target {
		if g.State == 1 {
			c++
		}
	}

	chars := float64(len(totalInput))

	return float64(c/chars) * 100
}
