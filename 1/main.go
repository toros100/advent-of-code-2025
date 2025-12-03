package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./1/input")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	dial := 50
	passedZeroes := 0
	exactZeroes := 0
	var sign int

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == 'L' {
			sign = -1
		} else {
			sign = 1
		}

		val, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		for range val {
			dial = (dial + sign) % 100
			if dial < 0 {
				dial += 100
			}
			if dial == 0 {
				passedZeroes++
			}
		}

		if dial == 0 {
			exactZeroes++
		}
	}
	fmt.Printf("Exact 0s (part 1): %d\nPassed 0s (part 2): %d\n", exactZeroes, passedZeroes)
}
