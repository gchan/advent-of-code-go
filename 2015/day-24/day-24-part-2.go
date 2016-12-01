package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func generateCombinations(elements []int, length int) <-chan []int {
	c := make(chan []int)
	combo := make([]int, length)
	var nextCombo func(int, int)
	last := len(combo) - 1

	nextCombo = func(i, from int) {
		for j := from; j < len(elements); j++ {
			combo[i] = elements[j]
			if i == last {
				c <- combo
			} else {
				nextCombo(i+1, j+1)
			}
		}
	}

	go func(c chan []int) {
		defer close(c)
		nextCombo(0, 0)
	}(c)

	return c
}

func main() {
	input, err := ioutil.ReadFile("./day-24-input.txt")
	if err != nil {
		panic(err)
	}

	presentsStrs := strings.Split(string(input), "\n")
	presents := make([]int, len(presentsStrs))

	for i, present := range presentsStrs {
		presents[i], _ = strconv.Atoi(present)
	}

	groups := 4

	sum := 0
	minQe := math.MaxInt64
	for _, present := range presents {
		sum += present
	}

	target := sum / groups
	maxGroupSize := len(presents) / groups

	for groupSize := 1; groupSize < maxGroupSize+1; groupSize++ {
		minEqFound := false
		combos := generateCombinations(presents, groupSize)

		for combo := range combos {
			sum := 0
			qe := 1

			for _, present := range combo {
				sum += present
				qe *= present
			}

			if sum == target && qe < minQe {
				minQe = qe
				minEqFound = true
			}
		}

		if minEqFound {
			break
		}
	}

	println(minQe)
}
