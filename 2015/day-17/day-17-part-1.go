package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./day-17-input.txt")
	if err != nil {
		panic(err)
	}

	descriptions := strings.Split(string(input), "\n")
	containers := []int{}
	liters := 150
	combinations := 0

	for _, description := range descriptions {
		size, _ := strconv.Atoi(description)
		containers = append(containers, size)
	}

	testCombo := func(combo []int) {
		sum := 0
		for _, container := range combo {
			sum += container
		}
		if sum == liters {
			combinations++
		}
	}

	for num := 1; num <= len(containers); num++ {
		combo := make([]int, num)

		var nextCombo func(int, int)
		last := len(combo) - 1
		nextCombo = func(i, from int) {
			for j := from; j < len(containers); j++ {
				combo[i] = containers[j]
				if i == last {
					testCombo(combo)
				} else {
					nextCombo(i+1, j+1)
				}
			}
		}

		nextCombo(0, 0)
	}

	println(combinations)
}
