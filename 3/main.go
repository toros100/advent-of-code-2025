package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./3/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sumOne, sumTwo := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		sumOne += maxJolt(line, 2)
		sumTwo += maxJolt(line, 12)
	}
	fmt.Printf("Maximal outputs\nPart 1: %d\nPart 2: %d\n", sumOne, sumTwo)
}

func maxJolt(line string, digits int) int {
	vs := make([]int, len(line))
	spl := strings.Split(line, "")
	for i, v := range spl {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		vs[i] = vInt
	}
	return maxJoltHelper(vs, digits)
}

func maxJoltHelper(vs []int, digitsLeft int) int {
	if digitsLeft > len(vs) {
		log.Fatal("not enough values (programmer error)")
	}
	if len(vs) == 0 {
		return 0
	}
	maxVal := 0
	maxIdx := 0
	for i, v := range vs {
		if v > maxVal && len(vs)-(i+1) >= digitsLeft-1 {
			maxVal = v
			maxIdx = i
		}
	}
	return maxVal*int(math.Pow(10.0, float64(digitsLeft-1))) + maxJoltHelper(vs[maxIdx+1:], digitsLeft-1)
}
