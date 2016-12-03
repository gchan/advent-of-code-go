package main

import (
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-02-input.txt")

	keys := [5][5]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}

	var pin string

	x := 0
	y := 2

	x2 := x
	y2 := y

	for _, sequence := range strings.Split(string(input), "\n") {
		for _, cmd := range sequence {
			switch cmd {
			case 'U':
				y2 = int(math.Max(float64(y)-1.0, 0))
			case 'D':
				y2 = int(math.Min(float64(y)+1.0, 4))
			case 'L':
				x2 = int(math.Max(float64(x)-1.0, 0))
			case 'R':
				x2 = int(math.Min(float64(x)+1.0, 4))
			}

			if keys[y2][x2] != "" {
				x = x2
				y = y2
			}

			x2 = x
			y2 = y
		}

		pin = pin + keys[y][x]
	}

	println(pin)
}
