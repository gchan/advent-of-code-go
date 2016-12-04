package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-03-input.txt")

	triangleStrings := strings.Split(string(input), "\n")
	var triangleCount int

	sides := make([]int, 3)

	for _, triangleString := range triangleStrings {
		for i, side := range strings.Fields(triangleString) {
			sides[i], _ = strconv.Atoi(side)
		}
		sort.Ints(sides)

		if sides[0]+sides[1] > sides[2] {
			triangleCount++
		}
	}

	println(triangleCount)
}
