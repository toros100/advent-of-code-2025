package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type exercise struct {
	op               int32
	colStart, colEnd int
	rowStart, rowEnd int
}
type worksheet struct {
	data [][]byte
	exs  []exercise
}

func newWorksheet(numberLines []string, opLine string) *worksheet {
	data := make([][]byte, len(numberLines))
	for i, line := range numberLines {
		data[i] = []byte(line)
	}

	exs := make([]exercise, 0)
	for i, c := range opLine {
		if c == ' ' {
			continue
		} else if c == '*' || c == '+' {
			exs = append(exs, exercise{op: c, rowStart: 0, rowEnd: len(numberLines), colStart: i, colEnd: -1})
		} else {
			log.Fatalf("invalid character on operator line: %c", c)
		}
	}
	for i := range exs {
		if i < len(exs)-1 {
			exs[i].colEnd = exs[i+1].colStart - 1
		} else {
			exs[i].colEnd = len(opLine)
		}
	}
	return &worksheet{
		data: data,
		exs:  exs,
	}
}

func (w *worksheet) calculateAll(transposed bool) int {
	res := 0
	for i, _ := range w.exs {
		res += w.calculate(i, transposed)
	}
	return res
}

func (w *worksheet) calculate(exIdx int, transposed bool) int {

	if exIdx < 0 || exIdx >= len(w.exs) {
		log.Fatalf("exIdx out of range: %d\n", exIdx)
	}
	e := w.exs[exIdx]

	var neutral int
	var opFunc func(int, int) int

	if e.op == '+' {
		neutral = 0
		opFunc = func(a, b int) int { return a + b }
	} else {
		neutral = 1
		opFunc = func(a, b int) int { return a * b }
	}

	res := neutral

	if transposed {
		b := make([]byte, e.rowEnd-e.rowStart)
		for i := e.colStart; i < e.colEnd; i++ {
			for j := e.rowStart; j < e.rowEnd; j++ {
				b[j] = w.data[j][i]
			}
			n, err := strconv.Atoi(strings.TrimSpace(string(b)))
			if err != nil {
				log.Fatalf("while parsing %s: %v\n", string(b), err)
			}
			res = opFunc(res, n)
		}
	} else {
		for i := e.rowStart; i < e.rowEnd; i++ {
			n, err := strconv.Atoi(strings.TrimSpace(string(w.data[i][e.colStart:e.colEnd])))
			if err != nil {
				log.Fatal(err)
			}
			res = opFunc(res, n)
		}
	}
	return res
}

func main() {
	f, err := os.Open("./6/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numberLines := make([]string, 0)
	var operatorLine string
	for scanner.Scan() {
		line := scanner.Text()
		_, err := strconv.Atoi(line[:1])
		if err == nil {
			// first char was a number, assuming number line
			numberLines = append(numberLines, line)
		} else {
			// first char was not number (+,*), assuming operator line
			operatorLine = line
			break
		}
	}

	sm := newWorksheet(numberLines, operatorLine)
	fmt.Printf("Part 1: %d\n", sm.calculateAll(false))
	fmt.Printf("Part 2: %d\n", sm.calculateAll(true))
}
