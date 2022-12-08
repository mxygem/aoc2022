package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	out, err := reorg(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result: %d\n", out)
}

func reorg(input []byte) (int, error) {
	var total int

	for _, g := range bytes.Split(input, []byte("\n")) {
		if len(g) == 0 {
			continue
		}

		var match byte
		for _, f := range g[:len(g)/2] {
			for _, b := range g[len(g)/2:] {
				if f != b {
					continue
				}
				match = b
				break
			}
			if match != 0 {
				break
			}
		}

		switch {
		case match >= 97:
			total += int(match) - 96
		case match <= 90:
			total += int(match) - 38
		}
	}

	return total, nil
}
