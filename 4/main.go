package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("./4/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	pd := newPrintingDepartment()

	lineNo := 0
	for scanner.Scan() {
		lineNo += 1
		line := scanner.Text()
		for i, ch := range line {
			if ch == '@' {
				pd.setPaper(pos{lineNo, i + 1})
			}
		}
		if pd.columns < len(line) {
			// can not go by size of map pd.grid[_] (which represents one column) due to empty struct tomfoolery
			// (not storing explicit false for no paper, implicit by absence of struct{})
			pd.columns = len(line)
		}
	}
	pd.lines = lineNo

	fmt.Printf("Paper accessible (part 1): %d\n", pd.collectPaper(false))
	fmt.Printf("Paper accessible with retry (part 2): %d\n", pd.collectPaper(true))

}

type pos struct{ l, c int }

func (p pos) getNeighbours() []pos {
	res := make([]pos, 8)
	res[0] = pos{p.l - 1, p.c - 1}
	res[1] = pos{p.l - 1, p.c}
	res[2] = pos{p.l - 1, p.c + 1}
	res[3] = pos{p.l, p.c - 1}
	res[4] = pos{p.l, p.c + 1}
	res[5] = pos{p.l + 1, p.c - 1}
	res[6] = pos{p.l + 1, p.c}
	res[7] = pos{p.l + 1, p.c + 1}
	return res
}

type printingDepartment struct {
	grid           map[int]map[int]struct{}
	lines, columns int
}

func newPrintingDepartment() *printingDepartment {
	return &printingDepartment{
		grid: make(map[int]map[int]struct{}),
	}
}

func (pd *printingDepartment) setPaper(p pos) {
	m, ok := pd.grid[p.l]
	if ok {
		m[p.c] = struct{}{}
	} else {
		pd.grid[p.l] = make(map[int]struct{})
		pd.grid[p.l][p.c] = struct{}{}
	}
}

func (pd *printingDepartment) isPaper(p pos) bool {
	_, ok := pd.grid[p.l][p.c]
	return ok
}

func (pd *printingDepartment) isAccessible(p pos) bool {
	n := p.getNeighbours()
	ps := 0
	for _, v := range n {
		if pd.isPaper(v) {
			ps++
			if ps > 3 {
				return false
			}
		}
	}
	return true
}

func (pd *printingDepartment) removePaper(p pos) {
	delete(pd.grid[p.l], p.c)
}

func (pd *printingDepartment) collectPaper(mutate bool) int {
	paperCollected := 0
	modified := false
	for {
		modified = false
		for line, m := range pd.grid {
			for col, _ := range m {
				p := pos{line, col}
				if pd.isPaper(p) && pd.isAccessible(p) {
					paperCollected++
					if mutate {
						pd.removePaper(p)
						modified = true
					}
				}
			}
		}
		if !modified {
			break
		}
	}
	return paperCollected
}
