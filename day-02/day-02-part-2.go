package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day-02-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		present := scanner.Text()
		sides := []int{}

		for _, sideString := range strings.Split(present, "x") {
			side, _ := strconv.Atoi(sideString)
			sides = append(sides, side)
		}

		sort.Ints(sides)

		// Wrap
		total += (sides[0] + sides[1]) * 2
		// Bow
		total += sides[0] * sides[1] * sides[2]
	}

	println(total)
}
