package practice

import (
	"math/rand"
	"time"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
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

func generateSessionStatsChart() linechart.Model {
	var lc linechart.Model

	// Chart dimensions in characters
	const chartWidth = 24
	const chartHeight = 8

	// World-coordinate range: tweak as you like
	const (
		minX = 0.0
		maxX = 10.0
		minY = 0.0
		maxY = 10.0
	)

	baseOpts := []linechart.Option{
		linechart.WithStyles(axisStyle, lineStyle, labelStyle),
		linechart.WithXYSteps(5, 5), // how many tick labels
		linechart.WithAutoXYRange(), // auto-expand if points go outside
	}

	lc = linechart.New(chartWidth, chartHeight, minX, maxX, minY, maxY, baseOpts...)

	lc.DrawXYAxisAndLabel()

	var randomFloat64Point1 canvas.Float64Point
	var randomFloat64Point2 canvas.Float64Point
	var randomFloat64Point3 canvas.Float64Point

	var randomFloat64Point4 canvas.Float64Point
	var randomFloat64Point5 canvas.Float64Point
	var randomFloat64Point6 canvas.Float64Point

	randomFloat64Point1 = canvas.Float64Point{X: 0, Y: 0}
	randomFloat64Point2 = canvas.Float64Point{X: 5, Y: 5}
	randomFloat64Point3 = canvas.Float64Point{X: 7, Y: 2}

	randomFloat64Point4 = canvas.Float64Point{X: 8, Y: 8}
	randomFloat64Point5 = canvas.Float64Point{X: 6, Y: 6}
	randomFloat64Point6 = canvas.Float64Point{X: 8, Y: 0}

	// linechart3 draws braille line
	lc.DrawBrailleLineWithStyle(randomFloat64Point1, randomFloat64Point2, lineStyle)

	// linechart3 draws braille line
	lc.DrawBrailleLineWithStyle(randomFloat64Point2, randomFloat64Point3, lineStyle)

	lc.DrawBrailleLineWithStyle(randomFloat64Point4, randomFloat64Point5, lineStyle)

	// linechart3 draws braille line
	lc.DrawBrailleLineWithStyle(randomFloat64Point5, randomFloat64Point6, lineStyle)

	return lc
}
