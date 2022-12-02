package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var input string
	flag.StringVar(&input, "input", "", "specify the location of the input file to process")
	flag.Parse()

	if input == "" {
		log.Fatal("no input file specified")
	}

	invs, err := readInput(input)
	if err != nil {
		log.Fatalf("failed to read input: %s", err)
	}

	elves := elfInventories(invs)
	if len(elves) == 0 {
		log.Fatal("no elves with valid inventories found")
	}

	log.Printf("found %d elves", len(elves))
	log.Printf("the elf with the most calories is elf #%d with %d", elves[0].origPos+1, elves[0].total)

	log.Println("the top three elves are:")
	top3 := elves[:3]
	for _, e := range top3 {
		log.Printf("elf #%d with %d", e.origPos+1, e.total)
	}
	log.Printf("and together, they are holding %d worth of calories", top3[0].total+top3[1].total+top3[2].total)
}

func readInput(loc string) ([]string, error) {
	f, err := os.Open(loc)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var o []string
	for s.Scan() {
		l := s.Text()
		o = append(o, l)
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("scanning: %w", err)
	}

	return o, nil
}

type elf struct {
	origPos int
	total   int
	inv     []string
}

type elves []*elf

func elfInventories(invs []string) elves {
	if len(invs) == 0 {
		return nil
	}

	var es elves
	var e *elf
	for i, m := range invs {
		// handles multiple "newlines" in a row
		if m == "" && e == nil {
			continue
		}

		// found the end of an elf.
		if m == "" {
			e.total = countCalories(e.inv)
			es = append(es, e)
			e = nil
			continue
		}

		// Need to add new elf to list.
		if e == nil {
			e = &elf{origPos: len(es), inv: []string{m}}
			continue
		}

		// in the middle of collecting an elf's calories.
		e.inv = append(e.inv, m)

		// make sure last elf is added to collection.
		if i == len(invs)-1 {
			e.total = countCalories(e.inv)
			es = append(es, e)
		}
	}

	// sort elves desc by total.
	sort.Slice(es, func(i, j int) bool { return es[i].total > es[j].total })

	return es
}

func countCalories(inv []string) int {
	var t int
	for _, em := range inv {
		c, err := strconv.ParseInt(em, 0, 0)
		if err != nil {
			continue
		}
		t += int(c)
	}

	return t
}
