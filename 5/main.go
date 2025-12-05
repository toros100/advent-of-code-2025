package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type interval struct {
	start, end int
}

func main() {

	f, err := os.Open("./5/input")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	things := make([]interval, 0)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		if line == "" {
			// finished parsing intervals
			break
		}

		r := strings.Split(line, "-")

		if len(r) != 2 {
			log.Fatalf("unexpected input on line %d: %s", lineNo, line)
		}

		l, errL := strconv.Atoi(r[0])
		u, errU := strconv.Atoi(r[1])
		if errL != nil || errU != nil {
			log.Fatalf("failed to parse ints on line %d: %s", lineNo, line)
		}

		things = append(things, interval{start: l, end: u})
	}

	slices.SortFunc(things, func(a, b interval) int {
		return cmp.Compare(a.start, b.start)
	})

	compacted := make([]interval, 1)
	compacted[0] = things[0]
	prevIdx := 0
	for i := 1; i < len(things); i++ {
		if things[i].start <= compacted[prevIdx].end+1 {
			if things[i].end > compacted[prevIdx].end {
				compacted[prevIdx].end = things[i].end
			}
		} else {
			compacted = append(compacted, things[i])
			prevIdx += 1
		}
	}

	freshIngredients := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to parse int on line %d: %s", lineNo, line)
		}

		idx, found := slices.BinarySearchFunc(compacted, val, func(i interval, v int) int {
			return cmp.Compare(i.start, v)
		})
		if found {
			// found interval v-[...], clearly v is included
			freshIngredients += 1
		} else {
			// now, idx is the smallest index j such that:
			// (1) val < compacted[j].start
			// (2) for all 0 < i < j, we have compacted[i].start < val
			// thus, val is in any interval iff idx > 0 and val <= compacted[idx-1].end
			if idx > 0 && val <= compacted[idx-1].end {
				freshIngredients += 1
			}
		}
	}

	fmt.Printf("Fresh available ingredients (part 1): %d\n", freshIngredients)

	freshIds := 0
	for _, r := range compacted {
		freshIds += 1 + r.end - r.start
	}
	fmt.Printf("Fresh ingredient IDs (part 2): %d\n", freshIds)
}
