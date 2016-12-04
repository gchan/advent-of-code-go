package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-03-input.txt")

	sides := strings.Fields(string(input))

	var triangleCount int

	triangle := make([]int, 3)

	for i := 0; i < len(sides)/3; i++ {
		start := (i % 3) + (i/3)*9

		triangle[0], _ = strconv.Atoi(sides[start])
		triangle[1], _ = strconv.Atoi(sides[start+3])
		triangle[2], _ = strconv.Atoi(sides[start+6])

		sort.Ints(triangle)

		if triangle[0]+triangle[1] > triangle[2] {
			triangleCount++
		}
	}

	println(triangleCount)
}
