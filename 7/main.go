package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type tachyonState struct {
	beams  map[int]int
	minIdx int
	maxIdx int
}

func newTachyonState(line string) *tachyonState {
	bs := []byte(line)
	beams := make(map[int]int)
	beamIdx := slices.Index(bs, 'S')
	if beamIdx == -1 {
		log.Fatal("no initial beam")
	}
	beams[beamIdx] = 1
	return &tachyonState{
		beams:  beams,
		minIdx: 0,
		maxIdx: len(bs),
	}
}

func (ts *tachyonState) apply(line string) int {
	bs := []byte(line)
	splits := 0
	for i := range bs {
		if bs[i] == '^' {
			splits += ts.splitBeamIfPresent(i)
		}
	}
	return splits
}

func (ts *tachyonState) splitBeamIfPresent(idx int) int {
	if _, ok := ts.beams[idx]; !ok {
		return 0
	}
	if ts.minIdx < idx {
		ts.beams[idx-1] += ts.beams[idx]
	}
	if ts.maxIdx > idx {
		ts.beams[idx+1] += ts.beams[idx]
	}
	delete(ts.beams, idx)

	return 1
}

func (ts *tachyonState) currentTimelines() int {
	timelines := 0
	for _, t := range ts.beams {
		timelines += t
	}
	return timelines
}

func main() {
	f, err := os.Open("./7/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalSplits := 0
	var ts *tachyonState
	for scanner.Scan() {
		if ts == nil { // on first line
			ts = newTachyonState(scanner.Text())
			continue
		}

		totalSplits += ts.apply(scanner.Text())
	}

	if ts != nil {
		fmt.Println("Number of splits (part 1): ", totalSplits)
		fmt.Println("Number of timelines (part 2): ", ts.currentTimelines())
	} else {
		log.Fatal("empty input")
	}
}
