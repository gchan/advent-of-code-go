package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-02-input.txt")

	keys := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var pin string

	x := 1
	y := 1

	for _, sequence := range strings.Split(string(input), "\n") {
		for _, cmd := range sequence {
			switch cmd {
			case 'U':
				y = int(math.Max(float64(y)-1.0, 0))
			case 'D':
				y = int(math.Min(float64(y)+1.0, 2))
			case 'L':
				x = int(math.Max(float64(x)-1.0, 0))
			case 'R':
				x = int(math.Min(float64(x)+1.0, 2))
			}
		}

		pin = pin + strconv.Itoa(keys[y][x])
	}

	println(pin)
}
