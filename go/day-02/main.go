package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	f, _ := os.ReadFile(os.Args[2])
	log.Println(score(f))
}

// point values
const (
	win      = 6
	draw     = 3
	lose     = 0
	rock     = 1
	paper    = 2
	scissors = 3
)

// letter values
const (
	A = 65 // rock, lose
	B = 66 // paper, draw
	C = 67 // scissors, win
	// the following aren't used since the second column
	// values are offset for ease of use.
	// X = 88
	// Y = 89
	// Z = 90
)

func score(input []byte) int {
	r := 0
	for _, g := range bytes.Split(input, []byte("\n")) {
		if len(g) == 0 {
			continue
		}
		r += scoreByOutcome(g[0], g[2]-23)
	}
	return r
}

func scoreByOutcome(them, outcome byte) int {
	switch them {
	case A: // rock
		switch outcome {
		case A: // lose
			return lose + scissors
		case B: // draw
			return draw + rock
		case C: // win
			return win + paper
		default:
			return 0
		}
	case B: // paper
		switch outcome {
		case A:
			return lose + rock
		case B:
			return draw + paper
		case C:
			return win + scissors
		default:
			return 0
		}
	case C: // scissors
		switch outcome {
		case A:
			return lose + paper
		case B:
			return draw + scissors
		case C:
			return win + rock
		default:
			return 0
		}
	default:
		return 0
	}
}
