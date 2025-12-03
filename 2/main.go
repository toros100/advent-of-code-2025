package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2/input")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sumOne, sumTwo := 0, 0
	scanner.Scan()
	line := scanner.Text()
	r := strings.Split(line, ",")

	for i := range r {
		bounds := strings.Split(r[i], "-")
		if len(bounds) != 2 {
			log.Fatalf("invalid input: %s", r)
		}

		lb, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatal(err)
		}

		ub, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatal(err)
		}

		for j := lb; j <= ub; j++ {
			jStr := strconv.Itoa(j)
			if isInvalidPartOne(jStr) {
				sumOne += j
			}
			if isInvalidPartTwo(jStr) {
				sumTwo += j
			}
		}
	}
	fmt.Printf("Sum of invalid IDs\nPart 1: %d\nPart 2: %d\n", sumOne, sumTwo)
}

func isInvalidHelper(id string, k int) bool {
	firstK := id[:k]
	if len(id)%k != 0 {
		return false
	}
	return strings.Repeat(firstK, len(id)/k) == id
}

func isInvalidPartOne(id string) bool {
	if len(id)%2 != 0 {
		return false
	}
	return isInvalidHelper(id, len(id)/2)
}

func isInvalidPartTwo(id string) bool {
	for i := len(id) / 2; i >= 1; i-- {
		if isInvalidHelper(id, i) {
			return true
		}
	}
	return false
}
