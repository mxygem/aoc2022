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

	out, err := reorgx3(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result: %d\n", out)
}

func reorgx3(input []byte) (int, error) {
	invs := bytes.Split(input, []byte("\n"))
	var total int

	var es [][]byte
	for i := 0; i < len(invs); i++ {
		c := invs[i]
		if len(c) == 0 {
			continue
		}

		if len(es) < 2 {
			es = append(es, c)
			continue
		}
		es = append(es, c)

		match := findMatches(es[0], es[1], es[2])

		switch {
		case match >= 97:
			total += int(match) - 96
		case match <= 90:
			total += int(match) - 38
		}
		es = nil
	}

	return total, nil
}

func findMatches(e1, e2, e3 []byte) byte {
	var pos byte
	for _, a := range e1 {
		for _, b := range e2 {
			if a != b {
				continue
			}
			pos = b

			if pos != 0 {
				for _, c := range e3 {
					if c != pos {
						continue
					}
					if c == pos {
						return pos
					}
				}
			}

		}
	}

	return 0
}
