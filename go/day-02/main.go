package main

import (
	"bytes"
	"log"
	"os"
)

const (
	win  = 6
	draw = 3
	lose = 0
)

const (
	rock     = 1 // 65, 88
	paper    = 2 // 66, 89
	scissors = 3 // 67, 90
)

func main() {
	f, _ := os.ReadFile(os.Args[2])
	r := 0
	for _, g := range bytes.Split(f, []byte("\n")) {
		if len(g) == 0 {
			continue
		}

		r += calcScore(g[0], g[2]-23)
	}
	log.Println(r)
}

func calcScore(x, y byte) int {
	switch x {
	case 65:
		switch y {
		case 65:
			return draw + rock
		case 66:
			return win + paper
		case 67:
			return lose + scissors
		default:
			return 0
		}
	case 66:
		switch y {
		case 65:
			return lose + rock
		case 66:
			return draw + paper
		case 67:
			return win + scissors
		default:
			return 0
		}
	case 67:
		switch y {
		case 65:
			return win + rock
		case 66:
			return lose + paper
		case 67:
			return draw + scissors
		default:
			return 0
		}
	default:
		return 0
	}
}
